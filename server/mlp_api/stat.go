package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) Stat(context.Context, *StatRequest) (*StatResponse, error) {

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	count := db.GetEventCount(tx)


	return &StatResponse{
		Version: s.version,
		EventCount:count,
	}, nil
}

