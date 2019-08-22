package mlp_api

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


type server struct{
	db *db.DB
	version string
}

func NewServer(db *db.DB, version string) CatalogServer{
	return &server{db, version}
}

func (c *server) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return &StatResponse{
		Version:c.version,
	}, nil
}

func (c *server) CreateProject(ctx context.Context, r *CreateProjectRequest) (*CreateProjectResponse, error) {
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
		Meta:r.Meta,
	}

	c.publish(tx, prj)
	tx.MustCommit()

	return &CreateProjectResponse{}, nil
}


func (c *server) publish(tx *db.Tx, e proto.Message) uint64{
	projection.Handle(tx, e)
	return db.AppendEvent(tx, e)
}





