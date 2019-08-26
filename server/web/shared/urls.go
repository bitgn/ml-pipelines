package shared

import (
	"fmt"
	"log"
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

	// sanity
	if len(project)==0{
		log.Panicln("Project name can't be nil")
	}
	if len(dataset) == 0 {
		log.Panicln("Dataset name can't be nil")
	}



	return fmt.Sprintf("/projects/%s/datasets/%s", project, dataset)
}


