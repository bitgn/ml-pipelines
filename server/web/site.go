package web

import (
	"fmt"
	"path"
)

type Site struct{
	ProjectCount int
	DatasetCount int
	JobCount int
	ModelCount int
	ReportCount int
	ExpertCount int
	AppVersion string
	ActiveMenu string
	Url UrlResolver
}

type UrlResolver struct {

}

func (b *UrlResolver) Static(resource string) string{
	return path.Join("static", resource)
}

func (b *UrlResolver) ListProjects() string{
	return "explore"
}

func (b *UrlResolver) ExploreDatasets() string{
	return ""
}

func (b *UrlResolver) ViewProject(id string) string{
	return fmt.Sprintf("/projects/%s/", id)
}

func (b *UrlResolver) ViewDataset(dataset_id string) string{
	return ""
}



