package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (c *server) CreateProject(ctx context.Context, r *CreateProjectRequest) (*CreateProjectResponse, error) {

	wrap := func (err *ApiError) (*CreateProjectResponse, error){
		return &CreateProjectResponse{
			Error:err,
		}, nil
	}



	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return wrap(badName(vo.ENTITY_PROJECT, r.Name, err))
	}


	tx := c.db.MustWrite()
	defer tx.MustCleanup()


	pid := db.LookupProject(tx, r.Name)
	if len(pid)>0{
		return wrap(alreadyExists(vo.ENTITY_PROJECT, r.Name, r.Name, pid))
	}


	pid = sim.NewID()

	prj := &events.ProjectCreated{
		Uid:  pid,
		Name: r.Name,
		Meta: r.Meta,
	}

	c.publish(tx, prj)
	tx.MustCommit()

	return &CreateProjectResponse{
		Uid:pid,
	}, nil
}

