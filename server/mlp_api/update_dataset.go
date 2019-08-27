package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/vo"
)

func (c *server) UpdateDataset(ctx context.Context, r *UpdateDatasetRequest) (*ApiResponse, error) {
	tx := c.db.MustWrite()
	defer tx.MustCleanup()

	ds := db.GetDataset(tx,r.Uid)

	if ds == nil {
		return genError(notFound(vo.ENTITY_DATASET, r.Uid))
	}

	// TODO publish only deltas

	c.publish(tx, &events.DatasetUpdated{
		Uid:        r.Uid,
		ProjectUid: ds.ProjectUid,
		Meta:       r.Meta,
	})

	tx.MustCommit()
	return &ApiResponse{}, nil
}
