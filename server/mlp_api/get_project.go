package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
)

func (s *server) GetProject(c context.Context, r *GetProjectRequest) (*ProjectInfoResponse, error) {


	wrap := func (err *ApiError) (*ProjectInfoResponse, error){
		return &ProjectInfoResponse{
			Error:err.Method("GetProject"),
		}, nil
	}

	tx := s.db.MustRead()
	defer tx.MustCleanup()


	uid := db.LookupProject(tx, r.Name)
	if uid == nil {
		return wrap(unknownProjectName(r.Name))
	}

	prj := db.GetProject(tx, uid)

	if prj == nil {
		log.Panicln("Project not found ", uid)
	}



	return &ProjectInfoResponse{
		Name:prj.Name,
		Uid:prj.Uid,
	}, nil
}
