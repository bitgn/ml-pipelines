package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/vo"
)

func (s *server) GetJob(c context.Context, req *GetJobRequest) (*JobInfoResponse, error) {
	wrap := func (err *ApiError) (*JobInfoResponse, error){
		return &JobInfoResponse{
			Error:err.Method("GetJob"),
		}, nil
	}


	tx := s.db.MustRead()
	defer tx.MustCleanup()


	prj := db.GetProject(tx, req.ProjectUid)

	if prj == nil{
		return wrap(notFound(vo.ENTITY_PROJECT, req.ProjectUid))
	}



	entity := db.Lookup(tx, req.ProjectUid, req.Name)

	if entity == nil || entity.Kind != vo.ENTITY_JOB {
		return wrap(nameNotFound(vo.ENTITY_JOB, "", req.Name))
	}


	job := db.GetJob(tx, entity.Uid)
	if job == nil {
		log.Panicln("Job not found")
	}
	return &JobInfoResponse{
		Uid:entity.Uid,
		Name:job.Name,
		ProjectName:prj.Name,
		ProjectUid:prj.Uid,
	}, nil
}
