package projection

import (
	"github.com/golang/protobuf/proto"
	"mlp/catalog/db"
	"mlp/catalog/events"
)


func Handle(tx *db.Tx, msg proto.Message){
	switch e := msg.(type) {

	case *events.ProjectCreated:
		db.AddProject(tx, &db.ProjectData{Id: e.ProjectId,Name: e.Name,})


		stats := db.GetStats(tx)
		stats.ProjectCount +=1
		db.SetStats(tx, stats)
	case *events.DatasetCreated:

		data := &db.DatasetData{Name: e.Name, ProjectId: e.ProjectId, DatasetId: e.DatasetId}
		mergeMetadata(e.Meta, data)
		db.AddDataset(tx, data)

		stats := db.GetStats(tx)
		stats.DatasetCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectId)
		prj.DatasetCount +=1
		db.AddProject(tx, prj)

	case *events.DatasetUpdated:
		ds := db.GetDataset(tx, e.DatasetId)
		mergeMetadata(e.Meta, ds)
		db.UpdDataset(tx, ds)

	case *events.JobAdded:
		stats := db.GetStats(tx)
		stats.JobCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectId)
		prj.JobCount +=1
		db.AddProject(tx, prj)


		for _, input_id := range e.Inputs{
			ds := db.GetDataset(tx, input_id)
			ds.DownstreamJobs = append(ds.DownstreamJobs, e.JobId)
			db.UpdDataset(tx, ds)
		}
		for _, output_id := range e.Outputs{
			ds := db.GetDataset(tx, output_id)
			ds.UpstreamJobs = append(ds.UpstreamJobs, e.JobId)
			db.UpdDataset(tx, ds)
		}

		db.PutJob(tx, &db.Job{
			JobId:e.JobId,
			JobName:e.JobName,
			Inputs:e.Inputs,
			Outputs:e.Outputs,
			ProjectId:e.ProjectId,
		})


	case *events.ExpertAdded:
		stats := db.GetStats(tx)
		stats.ExpertCount +=1
		db.SetStats(tx, stats)


	}


}

func mergeMetadata(src *events.DatasetMetadata, trg *db.DatasetData) {
	if src.RecordCountSet {
		trg.RecordCount = src.RecordCount
	}
	if src.FileCountSet {
		trg.FileCount = src.FileCount
	}
	if src.StorageBytesSet  {
		trg.StorageBytes = src.StorageBytes
	}
	if src.UpdateTimestampSet {
		trg.UpdateTimestamp = src.UpdateTimestamp
	}
	if src.DataFormatSet {
		trg.DataFormat = src.DataFormat
	}
	if src.DescriptionSet {
		trg.Description = src.Description
	}
	if src.LocationIdSet{
		trg.LocationId = src.LocationId
	}
	if src.LocationUriSet {
		trg.LocationUri = src.LocationUri
	}
	if src.ExpertsSet{
		trg.Experts = src.Experts
	}
	if src.SampleSet {
		trg.Sample = src.Sample
	}
}
