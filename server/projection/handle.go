package projection

import (
	"github.com/golang/protobuf/proto"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/vo"
)


func Handle(tx *db.Tx, msg proto.Message){
	switch e := msg.(type) {

	case *events.ProjectCreated:
		data := &db.ProjectData{
			Uid:  e.Uid,
			Name: e.Name,
		}
		mergeProjectMeta(e.Meta, data)

		db.PutProject(tx, data)
		db.NameProject(tx, e.Uid, e.Name)

		stats := db.GetStats(tx)
		stats.ProjectCount +=1
		db.SetStats(tx, stats)
	case *events.DatasetCreated:

		data := &db.DatasetData{
			ProjectUid:  e.ProjectUid,
			Uid:         e.Uid,
			Name:        e.Name,
			ProjectName: e.ProjectName,
		}
		mergeDatasetMeta(e.Meta, data)

		db.PutDataset(tx, data)

		db.AddDatasetToProject(tx, e.ProjectUid, e.Uid)
		db.NameDataset(tx, e.ProjectName, e.Name, e.Uid)

		stats := db.GetStats(tx)
		stats.DatasetCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectUid)
		prj.StorageBytes += data.StorageBytes


		prj.DatasetCount +=1
		db.PutProject(tx, prj)

	case *events.DatasetUpdated:
		ds := db.GetDataset(tx, e.Uid)
		mergeDatasetMeta(e.Meta, ds)
		db.PutDataset(tx, ds)

		// recompute storage
		var storage int64
		for _, ds := range db.ListDatasetsFromProject(tx, e.ProjectUid){
			storage += ds.StorageBytes
		}

		prj := db.GetProject(tx, e.ProjectUid)
		prj.StorageBytes = storage
		db.PutProject(tx, prj)


	case *events.JobAdded:
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

		db.PutJob(tx, &db.Job{
			Uid:        e.Uid,
			Title:      e.Meta.Title,
			Name:       e.Name,
			Inputs:     e.Meta.Inputs,
			Outputs:    e.Meta.Outputs,
			ProjectUid: e.ProjectUid,
		})


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
	if d.RecordCountSet {
		trg.RecordCount = d.RecordCount
	}
	if d.FileCountSet {
		trg.FileCount = d.FileCount
	}
	if d.StorageBytesSet  {
		trg.StorageBytes = d.StorageBytes
	}
	if d.UpdateTimestampSet {
		trg.UpdateTimestamp = d.UpdateTimestamp
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
