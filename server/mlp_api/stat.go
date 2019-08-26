package mlp_api

import "golang.org/x/net/context"

func (c *server) Stat(context.Context, *StatRequest) (*StatResponse, error) {
	return &StatResponse{
		Version:c.version,
	}, nil
}

