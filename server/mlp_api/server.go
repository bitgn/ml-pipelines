package mlp_api

import (
	"github.com/golang/protobuf/proto"
	"mlp/catalog/db"
	"mlp/catalog/projection"
)


type server struct{
	db *db.DB
	version string
}

func genError(err *ApiError) (*ApiResponse, error){
	return &ApiResponse{
		Error:err,
	}, nil
}



func NewServer(db *db.DB, version string) CatalogServer{
	return &server{db, version}
}



func (c *server) publish(tx *db.Tx, e proto.Message) uint64{
	projection.Handle(tx, e)
	return db.AppendEvent(tx, e)
}





