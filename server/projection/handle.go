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
		data := &db.ProjectData{Id: e.ProjectId}
		mergeProjectMeta(e.Meta, data)

		db.AddProject(tx, data)

		stats := db.GetStats(tx)
		stats.ProjectCount +=1
		db.SetStats(tx, stats)
	case *events.DatasetCreated:

		data := &db.DatasetData{ProjectId: e.ProjectId, DatasetId: e.DatasetId}
		mergeDatasetMeta(e.Meta, data)
		db.AddDataset(tx, data)

		stats := db.GetStats(tx)
		stats.DatasetCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectId)
		prj.StorageBytes += data.StorageBytes


		prj.DatasetCount +=1
		db.AddProject(tx, prj)

	case *events.DatasetUpdated:
		ds := db.GetDataset(tx, e.DatasetId)
		mergeDatasetMeta(e.Meta, ds)
		db.UpdDataset(tx, ds)

		// recompute storage
		var storage int64
		for _, ds := range db.ListDatasets(tx, e.ProjectId){
			storage += ds.StorageBytes
		}

		prj := db.GetProject(tx, e.ProjectId)
		prj.StorageBytes = storage
		db.AddProject(tx, prj)


	case *events.JobAdded:
		stats := db.GetStats(tx)
		stats.JobCount +=1
		db.SetStats(tx, stats)

		prj := db.GetProject(tx, e.ProjectId)
		prj.JobCount +=1
		db.AddProject(tx, prj)


		for _, inputId := range e.Inputs{
			ds := db.GetDataset(tx, inputId)
			ds.DownstreamJobs = append(ds.DownstreamJobs, e.JobId)
			db.UpdDataset(tx, ds)
		}
		for _, outputId := range e.Outputs{
			ds := db.GetDataset(tx, outputId)
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

func mergeProjectMeta(d *vo.ProjectMetadataDelta, trg *db.ProjectData){
	if d.NameSet{
		trg.Name = d.Name
	}
	if d.DescriptionSet{
		trg.Description = d.Description
	}

}

func mergeDatasetMeta(d *vo.DatasetMetadataDelta, trg *db.DatasetData) {
	if d.NameSet {
		trg.Name = d.Name
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
