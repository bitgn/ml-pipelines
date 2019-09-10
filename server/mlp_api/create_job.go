package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)



func (s *server) CreateJob(_ context.Context, r *CreateJobRequest) (*JobInfoResponse, error) {


	genError := func (err *ApiError) (*JobInfoResponse, error){
		return &JobInfoResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_JOB, r.Name, err))
	}


	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	prj := db.GetProject(tx,r.ProjectUid)
	if prj == nil {
		return genError(notFound(vo.ENTITY_PROJECT, r.ProjectUid))
	}


	exists := db.Lookup(tx, r.ProjectUid, r.Name)
	if exists != nil {
		return genError(alreadyExists(exists.Kind, prj.Name, r.Name, exists.Uid))
	}



	uid := sim.NewID()

	e := &events.JobAdded{
		Name:       r.Name,
		ProjectUid: r.ProjectUid,
		Uid:        uid,
		Meta:       r.Meta,
		ProjectName:prj.Name,
	}

	if e.Meta == nil {
		e.Meta = &vo.JobMetadataDelta{}
	}

	s.publish(tx, e)
	tx.MustCommit()


	return &JobInfoResponse{
		Uid:uid,
		Name:r.Name,
		ProjectName:prj.Name,
		ProjectUid:prj.Uid,
	}, nil


}

