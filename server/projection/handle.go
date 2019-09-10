package projection

import (
	"github.com/golang/protobuf/proto"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/vo"
)


func mustName(val string){
	if len(val)==0{
		log.Panicf("Name can't be nil")
	}
}
func must(i interface{}){
	if i == nil {
		log.Panicln("Value can't be null")
	}
}



func Handle(tx *db.Tx, msg proto.Message){
	switch e := msg.(type) {

	case *events.ProjectCreated:
		mustName(e.Name)

		data := &db.ProjectData{
			Uid:  e.Uid,
			Name: e.Name,
		}
		mergeProjectMeta(e.Meta, data)
		if len(data.Title) == 0{
			data.Title = data.Name
		}

		db.PutProject(tx, data)
		db.NameProject(tx, e.Uid, e.Name)

		stats := db.GetStats(tx)
		stats.ProjectCount +=1
		db.PutStats(tx, stats)

	case *events.DatasetCreated:

		mustName(e.Name)
		mustName(e.ProjectName)

		data := &db.DatasetData{
			ProjectUid:  e.ProjectUid,
			Uid:         e.Uid,
			Name:        e.Name,
			ProjectName: e.ProjectName,
		}
		mergeDatasetMeta(e.Meta, data)

		if len(data.Title) == 0{
			data.Title = data.Name
		}

		db.PutDataset(tx, data)



		db.AddDatasetToProject(tx, e.ProjectUid, e.Uid)
		db.Name(tx, e.ProjectUid, e.Name, vo.ENTITY_DATASET, e.Uid)

		stats := db.GetStats(tx)
		stats.DatasetCount +=1
		db.PutStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectUid)
		prj.StorageBytes += data.StorageBytes


		prj.DatasetCount +=1
		db.PutProject(tx, prj)

	case *events.DatasetUpdated:


		data := db.GetDataset(tx, e.Uid)
		mergeDatasetMeta(e.Meta, data)

		db.PutDataset(tx, data)

		// recompute storage
		var storage int64
		for _, ds := range db.ListDatasetsFromProject(tx, e.ProjectUid){
			storage += ds.StorageBytes
		}

		prj := db.GetProject(tx, e.ProjectUid)
		prj.StorageBytes = storage
		db.PutProject(tx, prj)

	case *events.DatasetVersionAdded:
		ver := &db.DatasetVersionData{
			DatasetUid:e.DatasetUid,
			ParentUid:e.ParentUid,
			Title:e.Title,
			Uid:e.Uid,
			Items:e.Items,
			Inputs:e.Inputs,
			VersionNum:e.VersionNum,
			Timestamp:e.Timestamp,

		}


		if !ver.HasParent(){
			db.DeleteDatasetHead(tx, e.DatasetUid)
		}

		// this version becomes the head
		for _, i := range e.Items{
			db.PutDatasetHeadItem(tx, e.DatasetUid, &db.DatasetItemData{
				Name:i.Name,
				StorageBytes:i.StorageBytes,
				RecordsCount:i.Records,
			})
		}

		for _,i := range e.Remove{
			db.DelDatasetHeadItem(tx, e.DatasetUid, i.Name)
		}


		var (
			storage_bytes int64
			record_count int64
			item_count int64

		)

		for _, x := range db.ListDatasetHeadItems(tx, e.DatasetUid){
			item_count+=1
			record_count += x.RecordsCount
			storage_bytes += x.StorageBytes
		}

		// for each input in this dataset
		for _,input :=range e.Inputs{
			switch input.Type {
			case vo.DatasetVerInput_JOB_RUN:
				run := db.GetJobRun(tx, input.Uid)
				run.Outputs = append(run.Outputs, &db.JobRunOutput{
					Type:db.JobRunOutput_DatasetVer,
					Uid:e.Uid,
				})
				db.PutJobRun(tx, run)
			default:
				log.Panicln("Unhandled input type")
			}
		}



		ver.RecordCount = record_count
		ver.StorageBytes = storage_bytes
		ver.ItemCount = item_count

		db.PutDatasetVersion(tx,ver)

		ds := db.GetDataset(tx, e.DatasetUid)
		ds.UpdateTimestamp = e.Timestamp
		ds.VersionCount +=1


		ds.RecordCount = ver.RecordCount
		ds.StorageBytes = ver.StorageBytes
		ds.FileCount = ver.ItemCount
		ds.HeadVersion = ver.Uid

		db.PutDataset(tx, ds)



		var project_storage int64

		for _, ds := range db.ListDatasetsFromProject(tx, ds.ProjectUid) {
			project_storage += ds.StorageBytes
		}

		proj := db.GetProject(tx, ds.ProjectUid)
		proj.StorageBytes = project_storage
		db.PutProject(tx, proj)


	case *events.JobRunFailed:
		run := db.GetJobRun(tx,e.Uid)
		run.Status = vo.JOB_STATUS_FAIL
		run.UpdateTimestamp = e.Timestamp
		db.PutJobRun(tx, run)


	case *events.JobRunCompleted:
		run := db.GetJobRun(tx,e.Uid)
		run.Status = vo.JOB_STATUS_SUCCESS
		run.UpdateTimestamp = e.Timestamp
		db.PutJobRun(tx, run)

		job := db.GetJob(tx, e.JobUid)
		job.LastSuccessUid = e.Uid
		db.PutJob(tx, job)

	case *events.JobRunStarted:

		run := &db.JobRunData{
			Uid:e.Uid,
			JobUid:e.JobUid,
			Title:e.Title,
			StartTimestamp:e.Timestamp,
			UpdateTimestamp:e.Timestamp,
			Inputs:e.Inputs,
			Status:vo.JOB_STATUS_RUNNING,
		}



		db.PutJobRun(tx, run)

		// for each run input
		// XXXX (output) --> (input) This JOB.RUN
		for _, input := range run.Inputs{
			switch input.Type {
			case vo.JobRunInput_DatasetVer:
				ds := db.GetDatasetVersion(tx, input.Uid)
				ds.Outputs = append(ds.Outputs, &vo.DatasetVerOutput{
					Type:vo.DatasetVerOutput_JOB_RUN,
					Uid:run.Uid,
				})
				db.PutDatasetVersion(tx, ds)
			case vo.JobRunInput_System:

				// System -> THIS
				db.PutSystemLink(tx, input.Uid, &db.SystemLink{
					ContainerUid:e.JobUid,
					// for the link, this is output
					Type:db.SystemLink_Output_JobRun,
					InstanceUid:e.Uid,
				})

				// nothing so far
				//
			default:
				log.Panicf("Unknown job run input %s\n",input.Type)
			}
		}

		job := db.GetJob(tx, run.JobUid)
		job.RunNum =e.RunNum
		job.RunUid =e.Uid


		db.PutJob(tx, job)


	case *events.SystemCreated:
		mustName(e.Name)
		mustName(e.ProjectName)

		data := &db.SystemData{
			Name:e.Name,
			Uid:e.Uid,
			Title:e.Meta.Title,
			LocationId:e.Meta.LocationId,
			Description:e.Meta.Description,
			ProjectUid:e.ProjectUid,
			Experts:e.Meta.Experts,
			LocationUri:e.Meta.LocationUri,
			ProjectName:e.ProjectName,
			Kind:e.Meta.Kind,
		}

		db.PutSystem(tx, data)
		db.Name(tx, e.ProjectUid, e.Name, vo.ENTITY_SYSTEM, e.Uid)

		db.AddSystemToProject(tx, e.ProjectUid, e.Uid)


		proj := db.GetProject(tx, e.ProjectUid)
		proj.SystemCount +=1
		db.PutProject(tx, proj)


		stats := db.GetStats(tx)
		stats.SystemCount +=1
		db.PutStats(tx, stats)

	case *events.SystemVersionAdded:
		svc := db.GetSystem(tx, e.SystemUid)
		svc.VersionNum = e.Num;
		svc.VersionUid =e.Uid;
		svc.VersionTimestamp = e.Timestamp
		db.PutSystem(tx, svc)

		ver := &db.SystemVersionData{
			Uid:e.Uid,
			VersionNum:e.Num,
			Inputs:e.Inputs,
			Outputs:e.Outputs,
			Title:e.Title,
			Timestamp:e.Timestamp,
			SystemUid:e.SystemUid,
		}

		db.PutSystemVersion(tx, ver)
		db.IndexSystemVersion(tx,ver.SystemUid, ver.Uid, ver.VersionNum)


		for _, input := range e.Inputs{
			switch input.Type {
			case vo.SystemVersionInput_System:
				// System XXXX -> THIS
				db.PutSystemLink(tx, input.Uid, &db.SystemLink{
					ContainerUid:e.SystemUid,
					// for the link, this is output
					Type:db.SystemLink_Output_SystemVer,
					InstanceUid:e.Uid,
				})
			case vo.SystemVersionInput_JobRun:


				// JobRUN XXXX -> THIS

				run := db.GetJobRun(tx, input.Uid)
				run.Outputs = append(run.Outputs, &db.JobRunOutput{
					Type: db.JobRunOutput_SystemVer,
					Uid:e.Uid,
				})
				db.PutJobRun(tx, run)

			default:
				panic(input.Type)
			}
		}

		for _, output := range e.Outputs{
			switch output.Type {
			case vo.SystemVersionOutput_System:
				// System THIS -> System XXXX
				db.PutSystemLink(tx, output.Uid, &db.SystemLink{
					ContainerUid:e.SystemUid,
					// for the link, this is output
					Type:db.SystemLink_Input_SystemVer,
					InstanceUid:e.Uid,
				})
			default:
				panic(output.Type)
			}
		}




	case *events.JobAdded:

		mustName(e.Name)
		mustName(e.ProjectName)

		stats := db.GetStats(tx)
		stats.JobCount +=1
		db.PutStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectUid)
		prj.JobCount +=1
		db.PutProject(tx, prj)





		data := db.Job{
			Uid:        e.Uid,
			Title:      e.Meta.Title,
			Name:       e.Name,
			ProjectUid: e.ProjectUid,
			ProjectName:e.ProjectName,
		}
		if len(data.Title) == 0{
			data.Title = data.Name
		}
		db.PutJob(tx, &data)

		db.AddJobToProject(tx, e.ProjectUid, e.Uid)


		db.Name(tx, e.ProjectUid, e.Name, vo.ENTITY_JOB, e.Uid)


	case *events.ExpertAdded:
		stats := db.GetStats(tx)
		stats.ExpertCount +=1
		db.PutStats(tx, stats)



	}






}

func mergeProjectMeta(d *vo.ProjectMetadataDelta, trg *db.ProjectData){
	if d.TitleSet{
		trg.Title = d.Title
	}
	if d.DescriptionSet{
		trg.Description = d.Description
	}

}

func mergeDatasetMeta(d *vo.DatasetMetadataDelta, trg *db.DatasetData) {
	if d.TitleSet {
		trg.Title = d.Title
	}
	if d.DataFormatSet {
		trg.DataFormat = d.DataFormat
	}
	if d.DescriptionSet {
		trg.Description = d.Description
	}
	if d.LocationIdSet{
		trg.LocationId = d.LocationId
	}
	if d.LocationUriSet {
		trg.LocationUri = d.LocationUri
	}
	if d.ExpertsSet{
		trg.Experts = d.Experts
	}
	if d.SampleSet {
		trg.Sample = d.Sample
	}
}
