package mlp_api

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
)

func (s *server) Reset(c context.Context, r *ResetRequest) (*EmptyResponse, error) {

	tx := s.db.MustWrite()

	defer tx.MustCleanup()

	if err := tx.Tx.Drop(tx.DB, false); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to empty the DB"))
	}
	tx.MustCommit()
	return &EmptyResponse{}, nil
}


