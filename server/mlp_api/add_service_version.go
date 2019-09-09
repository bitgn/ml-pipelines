package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
)

func (s *server) AddServiceVersion(c context.Context, r *AddServiceVersionRequest) (*AddServiceVersionResponse, error) {

	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	svc := db.GetService(tx, r.ServiceUid)
	if svc == nil {
		log.Panicln("Project not found")
	}
	prj := db.GetProject(tx, svc.ProjectUid)
	if prj == nil {
		log.Panicln("Project not found")
	}

	uid := sim.NewID()

	e := &events.ServiceVersionAdded{
		ProjectUid: svc.ProjectUid,
		Uid:        uid,
		ServiceUid: svc.Uid,
		Num:        svc.VersionNum + 1,
		Title:      r.Title,
		Outputs:    r.Outputs,
		Inputs:     r.Inputs,
		Timestamp: sim.Unix(),
	}

	s.publish(tx, e)
	tx.MustCommit()

	return &AddServiceVersionResponse{
		Uid: uid,
	}, nil
}

