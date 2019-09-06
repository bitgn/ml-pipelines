package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) CompleteJobRun(c context.Context, r *CompleteJobRunRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	resp := &EmptyResponse{}
	run := db.GetJobRun(tx, r.Uid)

	if nil == run {
		return resp.err(notFound(vo.ENTITY_JOB_RUN, r.Uid))
	}

	switch run.Status {
	case vo.JOB_STATUS_SUCCESS, vo.JOB_STATUS_FAIL:
		return resp.err(jobRunState(r.Uid, run.Status))
	}

	e := &events.JobRunCompleted{
		Uid:       run.Uid,
		Timestamp: sim.Unix(),
		JobUid:    run.JobUid,
	}
	s.publish(tx, e)

	tx.MustCommit()
	return resp, nil

}

