package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) StartJobRun(c context.Context, r *StartJobRunRequest) (*JobRunInfoResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &JobRunInfoResponse{}
	job := db.GetJob(tx, r.JobUid)

	if nil == job{
		return resp.err(notFound(vo.ENTITY_JOB, r.JobUid))
	}

	if r.Timestamp == 0{
		r.Timestamp =sim.Unix()
	}

	e := &events.JobRunStarted{
		Uid:sim.NewID(),
		Title:r.Title,
		JobUid:r.JobUid,
		Timestamp:r.Timestamp,
		Inputs:r.Inputs,
		RunNum:job.RunNum+1,
	}
	s.publish(tx, e)

	resp.ProjectUid = job.ProjectUid
	resp.Uid = e.Uid

	tx.MustCommit()
	return resp, nil

}
