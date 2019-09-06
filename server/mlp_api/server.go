package mlp_api

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
)


type server struct{
	db *db.DB
	version string
}


func (s *server) StartJobRun(c context.Context, r *StartJobRunRequest) (*JobRunInfoResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &JobRunInfoResponse{}
	job := db.GetJob(tx, r.JobUid)

	if nil == job{
		return resp.err(notFound(vo.ENTITY_JOB, r.JobUid))
	}

	e := &events.JobRunStarted{
		Uid:sim.NewID(),
		Title:r.Title,
		JobUid:r.JobUid,
		Timestamp:sim.Unix(),
		Inputs:r.Inputs,
	}
	s.publish(tx, e)

	resp.ProjectUid = job.ProjectUid
	resp.Uid = e.Uid

	tx.MustCommit()
	return resp, nil

}

func (s *server) LogJobRun(c context.Context, r *LogJobRunRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &EmptyResponse{}
	job := db.GetJobRun(tx, r.Uid)

	if nil == job{
		return resp.err(notFound(vo.ENTITY_JOB_RUN, r.Uid))
	}

	e := &events.JobRunLogged{
		Uid:sim.NewID(),
		Timestamp:sim.Unix(),
		JobUid:job.Uid,
		Details:r.Details,
		LogTitle:r.LogTitle,

	}
	s.publish(tx, e)


	tx.MustCommit()
	return resp, nil

}

func (s *server) FailJobRun(c context.Context, r *FailJobRunRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &EmptyResponse{}
	job := db.GetJobRun(tx, r.Uid)

	if nil == job{
		return resp.err(notFound(vo.ENTITY_JOB_RUN, r.Uid))
	}


	switch job.Status{
	case vo.JOB_STATUS_SUCCESS, vo.JOB_STATUS_FAIL:
		return resp.err(jobRunState(r.Uid, job.Status))
	}

	e := &events.JobRunFailed{
		Uid:sim.NewID(),
		Timestamp:sim.Unix(),
		JobUid:job.Uid,
		Details:r.Details,
		Message:r.Message,
	}
	s.publish(tx, e)


	tx.MustCommit()
	return resp, nil
}

func (s *server) CompleteJobRun(c context.Context, r *CompleteJobRunRequest) (*EmptyResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &EmptyResponse{}
	job := db.GetJobRun(tx, r.Uid)

	if nil == job{
		return resp.err(notFound(vo.ENTITY_JOB_RUN, r.Uid))
	}


	switch job.Status{
	case vo.JOB_STATUS_SUCCESS, vo.JOB_STATUS_FAIL:
		return resp.err(jobRunState(r.Uid, job.Status))
	}

	e := &events.JobRunCompleted{
		Uid:sim.NewID(),
		Timestamp:sim.Unix(),
		JobUid:job.Uid,
	}
	s.publish(tx, e)


	tx.MustCommit()
	return resp, nil

}

func (s *server) GetDataset(c context.Context, req *GetDatasetRequest) (*DatasetInfoResponse, error) {


	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &DatasetInfoResponse{}
	prj := db.GetProject(tx, req.ProjectUid)

	if prj == nil{
		return resp.err(notFound(vo.ENTITY_PROJECT, req.ProjectUid))
	}



	entity := db.Lookup(tx, req.ProjectUid, req.Name)

	if entity == nil || entity.Kind != vo.ENTITY_DATASET {
		return resp.err(nameNotFound(vo.ENTITY_DATASET, "", req.Name))
	}


	ds := db.GetDataset(tx, entity.Uid)
	if ds == nil {
		log.Panicln("Dataset not found")
	}
	return &DatasetInfoResponse{
		Uid:         entity.Uid,
		Name:        ds.Name,
		ProjectName: prj.Name,
		ProjectUid:  prj.Uid,
		Description:ds.Description,
		UpdateTimestamp:ds.UpdateTimestamp,
		LocationUri:ds.LocationUri,
		LocationId:ds.LocationId,
		DataFormat:ds.DataFormat,
	}, nil
}

func (s *server) GetLastDatasetVersion(c context.Context, req *GetLastDatasetVersionRequest) (*DatasetVersionResponse, error) {
	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &DatasetVersionResponse{}

	ds := db.GetDataset(tx, req.DatasetUid)

	resp.Uid = ds.HeadVersion
	if ds.HeadVersion == nil {
		resp.NoVersion = true
	}

	return resp,nil
}

func (s *server) Commit(c context.Context, r *CommitRequest) (*CommitResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &CommitResponse{}



	for _, changeSet := range r.Datasets{
		ds := db.GetDataset(tx, changeSet.DatasetUid)

		if !bytes.Equal(ds.HeadVersion, changeSet.ParentVersionUid) {
				return resp.err(versionMismatch(changeSet.ParentVersionUid, ds.HeadVersion))
		}


		e := &events.DatasetVersionAdded{
			DatasetUid:ds.Uid,
			ProjectUid:ds.ProjectUid,
			Title:r.Title,
			Uid:sim.NewID(),
			Timestamp:sim.Unix(),
			ProjectName:ds.ProjectName,
			ParentUid:ds.HeadVersion,
			Inputs:changeSet.Inputs,
			Items:changeSet.Items,
			CleanSlate:changeSet.CleanSlate,
		}

		if len(changeSet.Remove)>0{
			log.Panicln("Implement changeset removal")
		}

		s.publish(tx, e)

		resp.DatasetVersions = append(resp.DatasetVersions, e.Uid)
	}


	tx.MustCommit()
	return resp, nil

}

func genError(err *ApiError) (*EmptyResponse, error){
	return &EmptyResponse{
		Error:err,
	}, nil
}



func NewServer(db *db.DB, version string) CatalogServer{
	return &server{db, version}
}



func (s *server) publish(tx *db.Tx, e proto.Message) uint64{
	projection.Handle(tx, e)
	return db.AppendEvent(tx, e)
}





