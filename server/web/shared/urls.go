package shared

import (
	"encoding/hex"
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

func (b *UrlResolver) ViewDatasetVersion(project, dataset string, uid []byte) string {
	return fmt.Sprintf("/projects/%s/datasets/%s/ver/%s", project, dataset, hex.EncodeToString(uid))
}
func (b *UrlResolver) ViewService(project, service string) string {
	return fmt.Sprintf("/projects/%s/services/%s", project, service)
}

func (b *UrlResolver) ViewServiceVer(project, service string, version int32) string {
	return fmt.Sprintf("/projects/%s/services/%s/ver/%d", project, service, version)
}
func (b *UrlResolver) ViewJob(project, name string) string {
	return fmt.Sprintf("/projects/%s/jobs/%s", project, name)
}
func (b*UrlResolver) ViewJobRun(project, job string, num int32) string{
	return fmt.Sprintf("/projects/%s/jobs/%s/run/%d", project, job, num)
}


