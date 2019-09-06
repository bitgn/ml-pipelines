package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/vo"
)

func (s *server) GetDataset(c context.Context, req *GetDatasetRequest) (*DatasetInfoResponse, error) {

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &DatasetInfoResponse{}
	prj := db.GetProject(tx, req.ProjectUid)

	if prj == nil {
		return resp.err(notFound(vo.ENTITY_PROJECT, req.ProjectUid))
	}

	entity := db.Lookup(tx, req.ProjectUid, req.Name)

	if entity == nil || entity.Kind != vo.ENTITY_DATASET {
		return resp.err(nameNotFound(vo.ENTITY_DATASET, "", req.Name))
	}

	ds := db.GetDataset(tx, entity.Uid)
	if ds == nil {
		log.Panicln("Dataset not found")
	}
	return &DatasetInfoResponse{
		Uid:             entity.Uid,
		Name:            ds.Name,
		ProjectName:     prj.Name,
		ProjectUid:      prj.Uid,
		Description:     ds.Description,
		UpdateTimestamp: ds.UpdateTimestamp,
		LocationUri:     ds.LocationUri,
		LocationId:      ds.LocationId,
		DataFormat:      ds.DataFormat,
	}, nil
}

