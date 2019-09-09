package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/vo"
)

func (s *server) GetSystem(c context.Context, req *GetSystemRequest) (*SystemInfoResponse, error) {

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &SystemInfoResponse{}
	prj := db.GetProject(tx, req.ProjectUid)

	if prj == nil {
		return resp.err(notFound(vo.ENTITY_PROJECT, req.ProjectUid))
	}

	entity := db.Lookup(tx, req.ProjectUid, req.Name)

	if entity == nil || entity.Kind != vo.ENTITY_SYSTEM {
		return resp.err(nameNotFound(vo.ENTITY_SYSTEM, "", req.Name))
	}

	ds := db.GetSystem(tx, entity.Uid)
	if ds == nil {
		log.Panicln("System not found")
	}
	return &SystemInfoResponse{
		Uid:         entity.Uid,
		Name:        ds.Name,
		ProjectName: prj.Name,
		ProjectUid:  prj.Uid,
	}, nil
}

