package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) GetDataset(c context.Context, r *GetDatasetRequest) (*DatasetInfoResponse, error) {

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &DatasetInfoResponse{}
	ds := db.GetDataset(tx, r.ProjectId, r.DatasetId)

	if ds == nil {
		return resp.err(datasetNotFound(r.ProjectId, r.DatasetId))
	}

	return &DatasetInfoResponse{
		DatasetId:       r.DatasetId,
		ProjectId:       r.ProjectId,
		UpdateTimestamp: ds.UpdateTimestamp,
		Summary:         ds.Summary,
		Sample:ds.Sample,
		Description:ds.Description,
	}, nil
}

