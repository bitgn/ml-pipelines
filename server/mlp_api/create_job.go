package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (c *server) CreateJob(_ context.Context, r *CreateJobRequest) (*CreateJobResponse, error) {


	genError := func (err *ApiError) (*CreateJobResponse, error){
		return &CreateJobResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_JOB, r.Name, err))
	}


	tx := c.db.MustWrite()
	defer tx.MustCleanup()

	prj := db.GetProject(tx,r.ProjectUid)
	if prj == nil {
		return genError(notFound(vo.ENTITY_PROJECT, r.ProjectUid))
	}


	exists := db.Lookup(tx, prj.Name, r.Name)
	if exists != nil {
		return genError(alreadyExists(exists.Kind, prj.Name, r.Name, exists.Uid))
	}


	uid := sim.NewID()

	e := &events.JobAdded{
		Name:       r.Name,
		ProjectUid: r.ProjectUid,
		Uid:        uid,
		Meta:       r.Meta,
		ProjectName:r.Name,
	}

	c.publish(tx, e)
	tx.MustCommit()


	return &CreateJobResponse{
		Uid:uid,
	}, nil


}

