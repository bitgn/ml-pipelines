package mlp_api

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/projection"
)


type server struct{
	db *db.DB
	version string
}

func (s *server) GetProject(c context.Context, r *GetProjectRequest) (*ProjectInfoResponse, error) {


	wrap := func (err *ApiError) (*ProjectInfoResponse, error){
		return &ProjectInfoResponse{
			Error:err,
		}, nil
	}

	tx := s.db.MustRead()
	defer tx.MustCleanup()


	uid := db.LookupProject(tx, r.Name)
	if uid != nil {
		return wrap(unknownProjectName(r.Name))
	}

	prj := db.GetProject(tx, uid)

	if prj == nil {
		log.Panicln("Project not found ", uid)
	}



	return &ProjectInfoResponse{
		Name:prj.Name,
		Uid:prj.Uid,
	}, nil
}

func (s *server) StartJobRun(context.Context, *StartJobRunRequest) (*JobRunInfoResponse, error) {
	panic("implement me")
}

func (s *server) LogJobRun(context.Context, *LogJobRunRequest) (*EmptyResponse, error) {
	panic("implement me")
}

func (s *server) FailJobRun(context.Context, *FailJobRunRequest) (*EmptyResponse, error) {
	panic("implement me")
}

func (s *server) CompleteJobRun(context.Context, *CompleteJobRunRequest) (*EmptyResponse, error) {
	panic("implement me")
}

func (s *server) GetDataset(context.Context, *GetDatasetRequest) (*DatasetInfoResponse, error) {
	panic("implement me")
}

func (s *server) GetLastDatasetVersion(context.Context, *GetLastDatasetVersionRequest) (*DatasetVersionResponse, error) {
	panic("implement me")
}

func (s *server) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	panic("implement me")
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





