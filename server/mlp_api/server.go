package mlp_api

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/projection"
	"reflect"
	"regexp"
)


type server struct{
	db *db.DB
	version string
	debug *regexp.Regexp
}

func genError(err *ApiError) (*EmptyResponse, error){
	return &EmptyResponse{
		Error:err,
	}, nil
}



func NewServer(db *db.DB, version string, debug *regexp.Regexp) CatalogServer{
	return &server{db, version,debug}
}



func (s *server) publish(tx *db.Tx, e proto.Message) uint64{

	if s.debug!=nil {

		kind := reflect.TypeOf(e).String()



		if s.debug.MatchString(kind){
			log.Println(kind)

			m := jsonpb.Marshaler{}
			result, _ := m.MarshalToString(e)
			log.Println(result)
		}

	}


	projection.Handle(tx, e)
	return db.AppendEvent(tx, e)
}





