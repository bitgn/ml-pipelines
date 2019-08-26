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

func (b *UrlResolver) ViewProject(name string) string{
	return fmt.Sprintf("/projects/%s", name)
}

func (b *UrlResolver) ViewDataset(project, dataset string) string{
	return fmt.Sprintf("/projects/%s/datasets/%s", project, dataset)
}


