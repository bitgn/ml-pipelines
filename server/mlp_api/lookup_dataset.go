package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/vo"
)

func (s *server) LookupDataset(c context.Context, r *LookupDatasetRequest) (*LookupDatasetResponse, error) {
	wrap := func (err *ApiError) (*LookupDatasetResponse, error){
		return &LookupDatasetResponse{
			Error:err,
		}, nil
	}



	err := domain.GetProblemsWithName(r.ProjectName)
	if err != nil {
		return wrap(badName(vo.ENTITY_PROJECT, r.ProjectName, err))
	}

	err = domain.GetProblemsWithName(r.DatasetName)
	if err != nil {
		return wrap(badName(vo.ENTITY_DATASET, r.DatasetName, err))
	}


	tx := s.db.MustRead()
	defer tx.MustCleanup()

	prj := db.LookupProject(tx, r.ProjectName)
	if prj == nil {
		return wrap(unknownProjectName(r.ProjectName))
	}


	found := db.Lookup(tx, r.ProjectName, r.DatasetName)
	if found == nil || found.Kind != vo.ENTITY_DATASET{
		return wrap( nameNotFound(vo.ENTITY_DATASET, r.ProjectName, r.DatasetName))
	}

	ds := db.GetDataset(tx, found.Uid)

	return &LookupDatasetResponse{
		ProjectUid:ds.ProjectUid,
		DatasetUid:ds.Uid,
	}, nil

}
