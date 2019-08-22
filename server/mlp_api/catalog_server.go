package mlp_api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
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

	resp := &CreateProjectResponse{}
	defer tx.MustCleanup()

	err := domain.GetProblemsWithID(r.ProjectId)



	if err != nil {
		apiError := &ApiError{
			Message: "Invalid ProjectID",
			SubjectId: "project_id",
			Code:StatusCode_INVALID_ARGUMENT,
			Details:[]string{err.Error()},
		}
		resp.Error = apiError

		return resp, nil
	}

	project := db.GetProject(tx, r.ProjectId)
	if project != nil {
		return &CreateProjectResponse{
			Error:&ApiError{
				Code:StatusCode_ALREADY_EXISTS,
				Message:fmt.Sprintf("Project '%s' already exists", r.ProjectId),
				SubjectId:r.ProjectId,

			},
		},nil
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





