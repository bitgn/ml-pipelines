package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) CreateDataset(ctx context.Context, r *CreateDatasetRequest) (*CreateDatasetResponse, error) {


	genError := func (err *ApiError) (*CreateDatasetResponse, error){
		return &CreateDatasetResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_DATASET, r.Name, err))
	}



	tx := s.db.MustWrite()
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

	e := &events.DatasetCreated{
		Name:       r.Name,
		ProjectUid: r.ProjectUid,
		Uid:        uid,
		Meta:       r.Meta,
		ProjectName: prj.Name,
	}

	s.publish(tx, e)
	tx.MustCommit()


	return &CreateDatasetResponse{
		Uid:uid,
	}, nil


}
