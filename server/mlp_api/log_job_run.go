package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) LogJobRun(c context.Context, r *LogJobRunRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	resp := &EmptyResponse{}
	run := db.GetJobRun(tx, r.Uid)

	if nil == run {
		return resp.err(notFound(vo.ENTITY_JOB_RUN, r.Uid))
	}

	if r.Timestamp == 0{
		r.Timestamp = sim.Unix()
	}

	e := &events.JobRunLogged{
		Uid:       run.Uid,
		Timestamp: r.Timestamp,
		JobUid:    run.JobUid,
		Details:   r.Details,
		LogTitle:  r.LogTitle,
	}
	s.publish(tx, e)

	tx.MustCommit()
	return resp, nil

}

