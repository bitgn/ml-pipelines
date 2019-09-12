package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) GetLastDatasetVersion(c context.Context, req *GetLastDatasetVersionRequest) (*DatasetVersionResponse, error) {
	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &DatasetVersionResponse{}

	ds := db.GetDataset(tx, req.DatasetUid)

	resp.Uid = ds.HeadVersion
	if ds.HeadVersion == nil {
		resp.NoVersion = true
	}

	return resp, nil
}


