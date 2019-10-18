package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
)

func (s *server) UpdateDataset(ctx context.Context, r *UpdateDatasetRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	ds := db.GetDataset(tx,r.ProjectId, r.DatasetId)

	if ds == nil {
		return genError(datasetNotFound(r.ProjectId, r.DatasetId))
	}

	// TODO publish only deltas

	s.publish(tx, &events.DatasetUpdated{
		ProjectId:r.ProjectId,
		DatasetId:r.DatasetId,
		Meta:       r.Meta,
	})

	tx.MustCommit()
	return &EmptyResponse{}, nil
}
