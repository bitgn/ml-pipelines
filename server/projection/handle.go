package projection

import (
	"github.com/golang/protobuf/proto"
	"mlp/catalog/db"
	"mlp/catalog/events"
)

const(
	DEL = events.STATE_DELETE
	VAL = events.STATE_VALUE
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

	case *events.DatasetUpdated:
		ds := db.GetDataset(tx, e.DatasetId)
		mergeMetadata(e.Meta, ds)
		db.UpdDataset(tx, ds)

	case *events.JobAdded:
		stats := db.GetStats(tx)
		stats.JobCount +=1
		db.SetStats(tx, stats)

	case *events.ExpertAdded:
		stats := db.GetStats(tx)
		stats.ExpertCount +=1
		db.SetStats(tx, stats)


	}


}

func mergeMetadata(src *events.DatasetMetadata, trg *db.DatasetData) {
	if src.RecordCountState == VAL {
		trg.RecordCount = src.RecordCount
	} else if src.RecordCountState == DEL {
		trg.RecordCount = 0
	}
	if src.FileCountState == VAL {
		trg.FileCount = src.FileCount
	} else if src.FileCountState == DEL {
		trg.FileCount = 0
	}
	if src.StorageBytesState == VAL {
		trg.StorageBytes = src.StorageBytes
	} else if src.StorageBytesState == DEL {
		trg.StorageBytes = 0
	}
	if src.UpdateTimestampState == VAL {
		trg.UpdateTimestamp = src.UpdateTimestamp
	} else if src.UpdateTimestampState == DEL {
		trg.UpdateTimestamp = 0
	}
	if src.DataFormatState == VAL {
		trg.DataFormat = src.DataFormat
	} else if src.DataFormatState == DEL {
		trg.DataFormat = ""
	}
	if src.DescriptionState == VAL {
		trg.Description = src.Description
	} else if src.DescriptionState == DEL {
		trg.Description = ""
	}
	if src.LocationIdState == VAL {
		trg.LocationId = src.LocationId
	} else if src.LocationIdState == DEL {
		trg.LocationId = ""
	}
	if src.LocationUriState == VAL {
		trg.LocationUri = src.LocationUri
	} else if src.LocationUriState == DEL {
		trg.LocationUri = ""
	}
	if src.ExpertsState == VAL {
		trg.Experts = src.Experts
	} else if src.ExpertsState == DEL {
		trg.Experts = nil
	}
	if src.SampleState == VAL {
		trg.Sample = src.Sample
	} else if src.SampleState == DEL {
		trg.Sample = nil
	}
}
