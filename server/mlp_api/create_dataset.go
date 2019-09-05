package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) CreateDataset(ctx context.Context, r *CreateDatasetRequest) (*DatasetInfoResponse, error) {


	genError := func (err *ApiError) (*DatasetInfoResponse, error){
		return &DatasetInfoResponse{
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
		log.Panicln("Project not found")
	}


	exists := db.Lookup(tx, r.ProjectUid, r.Name)
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

	return &DatasetInfoResponse{
		Uid:uid,
		ProjectUid:r.ProjectUid,
		ProjectName:prj.Name,
		Name:r.Name,
	}, nil


}
