package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) ListDatasets(c context.Context, r *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	resp := &ListDatasetsResponse{}
	tx := s.db.MustRead()
	defer tx.MustCleanup()

	prj := db.GetProject(tx, r.ProjectId)
	if prj == nil {
		return resp.err(projectNotFound(r.ProjectId))
	}

	dss := db.ListDatasets(tx, r.ProjectId)

	for _, ds := range dss {
		resp.Datasets = append(resp.Datasets, &DatasetInfo{
			ProjectId:       ds.ProjectId,
			Summary:         ds.Summary,
			DatasetId:       ds.DatasetId,
			Description:     ds.Description,
			Sample:          ds.Sample,
			UpdateTimestamp: ds.UpdateTimestamp,
		})
	}

	return resp, nil
}

