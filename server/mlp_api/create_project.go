package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
)






func (s *server) CreateProject(ctx context.Context, r *CreateProjectRequest) (*ProjectInfoResponse, error) {

	wrap := func (err *ApiError) (*ProjectInfoResponse, error){
		return &ProjectInfoResponse{
			Error:err,
		}, nil
	}



	err := domain.GetProblemsWithName(r.ProjectId)
	if err != nil {
		return wrap(badName(r.ProjectId, err))
	}


	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	pid := db.GetProject(tx, r.ProjectId)
	if pid!=nil{
		return wrap(projectAlreadyExists(r.ProjectId))
	}


	prj := &events.ProjectCreated{
		ProjectId:r.ProjectId,
	}

	s.publish(tx, prj)
	tx.MustCommit()

	return &ProjectInfoResponse{
		ProjectId:r.ProjectId,
	}, nil
}

