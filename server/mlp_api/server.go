package mlp_api

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/projection"
)


type server struct{
	db *db.DB
	version string
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

func (s *server) GetDataset(c context.Context, req *GetDatasetRequest) (*DatasetInfoResponse, error) {
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





