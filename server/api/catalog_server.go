package api

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/projection"
)


type catalogServer struct{
	db *db.DB
}


func (c *catalogServer) publish(tx *db.Tx, e proto.Message){
	projection.Handle(tx, e)
	db.AppendEvent(tx, e)
}

func (c *catalogServer) CreateProject(ctx context.Context, r *CreateProjectRequest) (*CreateProjectResponse, error) {
	tx := c.db.MustWrite()
	defer tx.MustCleanup()

	err := domain.GetProblemsWithID(r.ProjectId)

	if err != nil {
		return nil, err
	}

	project := db.GetProject(tx, r.ProjectId)
	if project != nil {
		return nil, errors.Errorf("Project '%s' already exists", r.ProjectId)
	}

	prj := &events.ProjectCreated{
		ProjectId:r.ProjectId,
		Name:r.ProjectName,
	}

	c.publish(tx, prj)
	tx.MustCommit()

	return &CreateProjectResponse{}, nil
}

func NewCatalogServer(db *db.DB) CatalogServer{
	return &catalogServer{db}
}



