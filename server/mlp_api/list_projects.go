package mlp_api

import (
	"golang.org/x/net/context"
	"mlp/catalog/db"
)

func (s *server) ListProjects(c context.Context, r *ListProjectsRequest) (*ListProjectsResponse, error) {
	resp := &ListProjectsResponse{}

	tx := s.db.MustRead()
	defer tx.MustCleanup()

	for _, prj := range db.ListProjects(tx) {
		resp.Projects = append(resp.Projects, &ProjectInfo{
			ProjectId: prj.ProjectId,
		})
	}

	return resp, nil
}

