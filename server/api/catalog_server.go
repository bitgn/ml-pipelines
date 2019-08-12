package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		st := status.New(codes.InvalidArgument,"invalid ProjectID")

		return nil, st.Err()
	}

	project := db.GetProject(tx, r.ProjectId)
	if project != nil {
		st := status.New(codes.AlreadyExists, fmt.Sprintf("Project '%s' already exists", r.ProjectId))
		return nil, st.Err()
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



