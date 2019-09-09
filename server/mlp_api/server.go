package mlp_api

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/events"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
	"mlp/catalog/vo"
	"reflect"
	"regexp"
)


type server struct{
	db *db.DB
	version string
	debug *regexp.Regexp
}

func (s *server) CreateService(c context.Context, r *CreateServiceRequest) (*ServiceInfoResponse, error) {

	genError := func (err *ApiError) (*ServiceInfoResponse, error){
		return &ServiceInfoResponse{
			Error:err,
		}, nil
	}


	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return genError(badName(vo.ENTITY_SERVICE, r.Name, err))
	}



	tx := s.db.MustWrite()
	defer tx.MustCleanup()



	prj := db.GetProject(tx,r.ProjectUid)
	if prj == nil {
		log.Panicln("Project not found")
	}


	exists := db.Lookup(tx, r.ProjectUid, r.Name)
	if exists != nil {
		return genError(alreadyExists(exists.Kind, prj.Name, r.Name, exists.Uid))
	}


	uid := sim.NewID()

	e := &events.ServiceCreated{
		Name:       r.Name,
		ProjectUid: r.ProjectUid,
		Uid:        uid,
		Meta:       r.Meta,
		ProjectName: prj.Name,
	}

	s.publish(tx, e)
	tx.MustCommit()

	return &ServiceInfoResponse{
		Uid:uid,
		ProjectUid:r.ProjectUid,
		ProjectName:prj.Name,
		Name:r.Name,
	}, nil


}

func (s *server) GetService(c context.Context, req *GetServiceRequest) (*ServiceInfoResponse, error) {

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	resp := &ServiceInfoResponse{}
	prj := db.GetProject(tx, req.ProjectUid)

	if prj == nil {
		return resp.err(notFound(vo.ENTITY_PROJECT, req.ProjectUid))
	}

	entity := db.Lookup(tx, req.ProjectUid, req.Name)

	if entity == nil || entity.Kind != vo.ENTITY_SERVICE {
		return resp.err(nameNotFound(vo.ENTITY_SERVICE, "", req.Name))
	}

	ds := db.GetService(tx, entity.Uid)
	if ds == nil {
		log.Panicln("Service not found")
	}
	return &ServiceInfoResponse{
		Uid:             entity.Uid,
		Name:            ds.Name,
		ProjectName:     prj.Name,
		ProjectUid:      prj.Uid,
	}, nil
}

func (s *server) AddServiceVersion(c context.Context, r *AddServiceVersionRequest) (*AddServiceVersionResponse, error) {


	tx := s.db.MustWrite()
	defer tx.MustCleanup()

	svc := db.GetService(tx, r.ServiceUid)
	if svc == nil {
		log.Panicln("Project not found")
	}
	prj := db.GetProject(tx,svc.ProjectUid)
	if prj == nil {
		log.Panicln("Project not found")
	}


	uid := sim.NewID()

	e := &events.ServiceVersionAdded{
		ProjectUid: svc.ProjectUid,
		Uid:        uid,
		ServiceUid:svc.Uid,
		Num:svc.VersionNum+1,
		Title:r.Title,
		Outputs:r.Outputs,
		Inputs:r.Inputs,


	}

	s.publish(tx, e)
	tx.MustCommit()

	return &AddServiceVersionResponse{
		Uid:uid,
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





