package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/vo"
)

func (s *server) LookupProject(ctx context.Context, r *LookupProjectRequest) (*LookupProjectResponse, error) {


	wrap := func (err *ApiError) (*LookupProjectResponse, error){
		return &LookupProjectResponse{
			Error:err,
		}, nil
	}



	err := domain.GetProblemsWithName(r.Name)
	if err != nil {
		return wrap(badName(vo.ENTITY_PROJECT, r.Name, err))
	}

	tx := s.db.MustRead()
	defer tx.MustCleanup()


	pid := db.LookupProject(tx, r.Name)
	if len(pid)==0{
		return wrap( unknownProjectName(r.Name))
	}

	return &LookupProjectResponse{
		Uid:pid,
	}, nil


}

