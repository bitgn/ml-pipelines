package mlp_api

import "golang.org/x/net/context"

func (s *server) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return &StatResponse{
		Version: s.version,
	}, nil
}

