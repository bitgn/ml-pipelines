package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/vo"
)

func (c *server) CreateDataset(ctx context.Context, r *CreateDatasetRequest) (*CreateDatasetResponse, error) {


	genError := func (err *ApiError) (*CreateDatasetResponse, error){
		return &CreateDatasetResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_DATASET, r.Name, err))
	}


	tx := c.db.MustWrite()
	defer tx.MustCleanup()





	prj := db.GetProject(tx,r.ProjectUid)
	if prj == nil {
		return genError(errNotFound(vo.ENTITY_PROJECT, r.ProjectUid))
	}


	exists := db.LookupDataset(tx, prj.Name, r.Name)
	if exists != nil {
		return genError(alreadyExists(vo.ENTITY_DATASET, r.Name, exists))
	}


	uid := newID()

	e := &events.DatasetCreated{
		Name:       r.Name,
		ProjectUid: r.ProjectUid,
		Uid:        uid,
		Meta:       r.Meta,
	}

	c.publish(tx, e)
	tx.MustCommit()


	return &CreateDatasetResponse{
		Uid:uid,
	}, nil


}
