package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/events"
	"mlp/catalog/sim"
)

func (s *server) AddDatasetActivity(ctx context.Context, r *AddDatasetActivityRequest) (*AddDatasetActivityResponse, error) {

	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	if r.Timestamp == 0 {
		r.Timestamp = sim.Unix()
	}

	e := &events.DatasetActivityAdded{
		UpdateTimestamp: r.Timestamp,
		DatasetId:       r.DatasetId,
		ProjectId:       r.ProjectId,
		Level:           r.Level,
		MultilineText:   r.MultilineText,
	}

	s.publish(tx, e)

	tx.MustCommit()

	return &AddDatasetActivityResponse{}, nil

}

