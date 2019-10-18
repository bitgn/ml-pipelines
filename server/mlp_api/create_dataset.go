package mlp_api

import (
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
)

func (s *server) CreateDataset(ctx context.Context, r *CreateDatasetRequest) (*DatasetInfoResponse, error) {


	genError := func (err *ApiError) (*DatasetInfoResponse, error){
		return &DatasetInfoResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.DatasetId)
	if err != nil {
		return genError(badName(r.DatasetId, err))
	}



	tx := s.db.MustWrite()
	defer tx.MustCleanup()



	prj := db.GetProject(tx,r.ProjectId)
	if prj == nil {
		log.Panicln("Project not found")
	}


	exists := db.GetDataset(tx, r.ProjectId, r.DatasetId)
	if exists != nil {
		return genError(datasetAlreadyExists(r.ProjectId, r.DatasetId))
	}




	e := &events.DatasetCreated{
		DatasetId:r.DatasetId,
		ProjectId:r.ProjectId,
		Meta:       r.Meta,
	}

	s.publish(tx, e)



	res := db.GetDataset(tx, r.ProjectId, r.DatasetId)

	tx.MustCommit()

	return &DatasetInfoResponse{
		DatasetId:       r.DatasetId,
		ProjectId:       r.ProjectId,
		UpdateTimestamp: res.UpdateTimestamp,
		Summary:         res.Summary,
		Sample:res.Sample,
		Description:res.Description,



	}, nil


}
