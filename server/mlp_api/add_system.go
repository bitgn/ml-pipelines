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

func (s *server) AddSystem(c context.Context, r *AddSystemRequest) (*SystemInfoResponse, error) {

	genError := func(err *ApiError) (*SystemInfoResponse, error) {
		return &SystemInfoResponse{
			Error: err,
		}, nil
	}

	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_SYSTEM, r.Name, err))
	}

	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	prj := db.GetProject(tx, r.ProjectUid)
	if prj == nil {
		log.Panicln("Project not found")
	}

	exists := db.Lookup(tx, r.ProjectUid, r.Name)
	if exists != nil {
		return genError(alreadyExists(exists.Kind, prj.Name, r.Name, exists.Uid))
	}

	uid := sim.NewID()

	e := &events.SystemCreated{
		Name:        r.Name,
		ProjectUid:  r.ProjectUid,
		Uid:         uid,
		Meta:        r.Meta,
		ProjectName: prj.Name,
	}

	s.publish(tx, e)
	tx.MustCommit()

	return &SystemInfoResponse{
		Uid:         uid,
		ProjectUid:  r.ProjectUid,
		ProjectName: prj.Name,
		Name:        r.Name,
	}, nil

}
