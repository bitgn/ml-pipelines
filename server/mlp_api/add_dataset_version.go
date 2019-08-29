package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)

func (s *server) AddDatasetVersion(c context.Context, r *AddDatasetVersionRequest) (*AddDatasetVersionResponse, error) {

	wrap := func (err *ApiError) (*AddDatasetVersionResponse, error){
		return &AddDatasetVersionResponse{
			Error:err,
		}, nil
	}



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

	if len(r.ParentUid) != 0 {
		found := db.GetDatasetVersion(tx, r.DatasetUid, r.ParentUid)
		if found == nil {
			return wrap(notFound(vo.ENTITY_DATASET_VERSION, r.ParentUid))
		}
	}

	e := &events.DatasetVersionAdded{
		Uid:sim.NewID(),

		ProjectUid:ds.ProjectUid,
		ProjectName:ds.ProjectName,
		Title:r.Title,
		ParentUid:r.ParentUid,
		Items:r.Items,
	}

	if r.Timestamp == 0{
		e.Timestamp = sim.Unix()
	}

	s.publish(tx, e)

	tx.MustCommit()

	return &AddDatasetVersionResponse{
		Uid:e.Uid,
	}, nil
}

