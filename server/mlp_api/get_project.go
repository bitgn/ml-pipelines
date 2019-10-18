package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) GetProject(c context.Context, r *GetProjectRequest) (*ProjectInfoResponse, error) {


	wrap := func (err *ApiError) (*ProjectInfoResponse, error){
		return &ProjectInfoResponse{}, nil
	}

	tx := s.db.MustRead()
	defer tx.MustCleanup()


	prj := db.GetProject(tx, r.ProjectId)
	if prj == nil {
		return wrap(projectNotFound(r.ProjectId))
	}




	return &ProjectInfoResponse{
		ProjectId:prj.ProjectId,
	}, nil
}
