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
		db.SetStats(tx, stats)

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
		db.Name(tx, e.ProjectName, e.Name, vo.ENTITY_DATASET, e.Uid)

		stats := db.GetStats(tx)
		stats.DatasetCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectUid)
		prj.StorageBytes += data.StorageBytes


		prj.DatasetCount +=1
		db.PutProject(tx, prj)

	case *events.DatasetUpdated:


		data := db.GetDataset(tx, e.Uid)
		mergeDatasetMeta(e.Meta, data)

		if len(data.Title) == 0{
			data.Title = data.Name
		}
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
		}

		for _, i := range e.Items{
			ver.StorageBytes += i.StorageBytes
			ver.RecordCount += i.Records
		}

		db.PutDatasetVersion(tx,ver)

		ds := db.GetDataset(tx, e.DatasetUid)
		ds.UpdateTimestamp = e.Timestamp

		if ver.HasParent() {
			// TODO: count total of all versions
			ds.RecordCount = ver.RecordCount
			ds.StorageBytes = ver.StorageBytes
			ds.FileCount = int32(len(ver.Items))



		} else {
			ds.RecordCount = ver.RecordCount
			ds.StorageBytes = ver.StorageBytes
			ds.FileCount = int32(len(ver.Items))
		}
		db.PutDataset(tx, ds)


	case *events.JobAdded:

		mustName(e.Name)
		mustName(e.ProjectName)

		stats := db.GetStats(tx)
		stats.JobCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectUid)
		prj.JobCount +=1
		db.PutProject(tx, prj)


		for _, input := range e.Meta.Inputs{

			switch input.Type {
			case vo.JobInput_Dataset:
				ds := db.GetDataset(tx, input.SourceId)
				ds.DownstreamJobs = append(ds.DownstreamJobs, e.Uid)
				db.PutDataset(tx, ds)
			}
		}
		for _, output := range e.Meta.Outputs{
			switch output.Type {
			case vo.JobOutput_Dataset:
				ds := db.GetDataset(tx, output.TargetId)
				ds.UpstreamJobs = append(ds.UpstreamJobs, e.Uid)
				db.PutDataset(tx, ds)

			}
		}



		data := db.Job{
			Uid:        e.Uid,
			Title:      e.Meta.Title,
			Name:       e.Name,
			Inputs:     e.Meta.Inputs,
			Outputs:    e.Meta.Outputs,
			ProjectUid: e.ProjectUid,
		}
		if len(data.Title) == 0{
			data.Title = data.Name
		}
		db.PutJob(tx, &data)


		db.Name(tx, e.ProjectName, e.Name, vo.ENTITY_JOB, e.Uid)


	case *events.ExpertAdded:
		stats := db.GetStats(tx)
		stats.ExpertCount +=1
		db.SetStats(tx, stats)
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
