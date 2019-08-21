package shared

import (
	"fmt"
	"path"
)

type UrlResolver struct {

}

func (b *UrlResolver) Static(resource string) string{
	return path.Join("/static", resource)
}

func (b *UrlResolver) ListProjects() string{
	return "/"
}

func (b *UrlResolver) ExploreDatasets() string{
	return "/explore"
}

func (b *UrlResolver) ViewProject(id string) string{
	return fmt.Sprintf("/projects/%s", id)
}

func (b *UrlResolver) ViewDataset(dataset_id string) string{
	return fmt.Sprintf("/datasets/%s", dataset_id)
}


