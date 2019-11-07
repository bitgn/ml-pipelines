package projection

import (
	"github.com/golang/protobuf/proto"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)


func mustName(val string){
	if len(val)==0{
		log.Panicf("Name can't be nil")
	}
}



func Handle(tx *db.Tx, msg proto.Message){
	switch e := msg.(type) {

	case *events.ProjectCreated:
		mustName(e.ProjectId)

		data := &db.ProjectData{
			ProjectId:e.ProjectId,

		}

		db.PutProject(tx, data)


	case *events.DatasetCreated:

		mustName(e.ProjectId)
		mustName(e.DatasetId)

		data := &db.DatasetData{
			DatasetId:e.DatasetId,
			ProjectId:e.ProjectId,
			UpdateTimestamp:sim.Unix(),
		}
		mergeDatasetMeta(e.Meta, data)
		db.PutDataset(tx, data)

	case *events.DatasetUpdated:
		mustName(e.ProjectId)
		mustName(e.DatasetId)

		data := db.GetDataset(tx, e.ProjectId, e.DatasetId)
		mergeDatasetMeta(e.Meta, data)
		db.PutDataset(tx, data)
	case *events.DatasetActivityAdded:
		mustName(e.ProjectId)
		mustName(e.DatasetId)

		data := &db.DatasetActivity{
			DatasetId:e.DatasetId,
			ProjectId:e.ProjectId,
			MultilineText:e.MultilineText,
			Level:e.Level,
			UpdateTimestamp:e.UpdateTimestamp,
			Id:db.NextActivityCounter(tx),
		}

		db.AddDatasetActivity(tx, data)

		ds := db.GetDataset(tx, e.ProjectId, e.DatasetId)

		if e.UpdateTimestamp > ds.UpdateTimestamp {
			ds.UpdateTimestamp = e.UpdateTimestamp
		}

		switch data.Level {
		case vo.ACTIVITY_LEVEL_ACTIVITY_ERROR:
			db.AddGlobalProblem(tx, data)
			db.AddGlobalActivity(tx, data)


			ds.Status = vo.DATASET_STATUS_STATUS_ERROR

		case vo.ACTIVITY_LEVEL_ACTIVITY_SUCCESS:
			db.AddGlobalActivity(tx, data)

			ds.Status = vo.DATASET_STATUS_STATUS_SUCCESS
		}


		db.PutDataset(tx, ds)
	}

}


func mergeDatasetMeta(d *vo.DatasetMetadataDelta, trg *db.DatasetData) {
	if d.DescriptionSet {
		trg.Description = d.Description
	}
	if d.SampleSet {
		trg.Sample = d.Sample
	}

	if d.SummarySet{
		trg.Summary = d.Summary
	}

}
