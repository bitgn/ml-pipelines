package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
)

func (s *server) AddDatasetVersion(c context.Context, r *AddDatasetVersionRequest) (*AddDatasetVersionResponse, error) {

	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	ds := db.GetDataset(tx, r.DatasetUid)


	//inputs := make([]vo.DatasetVerInput, len(r.Inputs))
	/*
	for _, input := range r.Inputs{

		dsi := db.GetDataset(tx, input)
		inputs = append(inputs, &vo.DatasetVerInput{
			Uid:input,
			Type:vo.DatasetVerInput_JOB_RUN,
		})

	}

	 */

	e := &events.DatasetVersionAdded{
		Uid:sim.NewID(),
		Timestamp:sim.Unix(),
		ProjectUid:ds.ProjectUid,
		ProjectName:ds.ProjectName,
		Title:r.Title,
		ParentUid:r.ParentUid,
		Items:r.Items,

	}

	s.publish(tx, e)

	tx.MustCommit()

	return &AddDatasetVersionResponse{
		Uid:e.Uid,
	}, nil
}

