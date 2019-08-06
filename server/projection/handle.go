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
	case *events.DatasetCreated:
		db.AddDataset(tx, &db.DatasetData{Name: e.Name,ProjectId: e.ProjectId,DatasetId: e.DatasetId})

	}

}
