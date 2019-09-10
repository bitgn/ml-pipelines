package view_project

import (
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
	"net/http"
)


type Dataset struct {
	Item *db.DatasetData
	IsStale bool
}



type JobItem struct {
	Item *db.Job

	CurrentStatus vo.JOB_STATUS
	LastUpdated int64


}

type Model struct {
	*shared.Site
	Project *db.ProjectData
	Datasets []*Dataset
	Systems []*db.SystemData
	Jobs []*JobItem
}


type Handler struct {
	Layout *shared.LoadedTemplate
	Env    *db.DB
}



func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		Layout: tl.DefineTemplate("web/layout.html","web/view_project/content.html"),
		Env:    env,
	}
}


func (h *Handler) Handle(w http.ResponseWriter, name string){
	tx := h.Env.MustRead()

	defer tx.MustAbort()


	pid := db.LookupProject(tx, name)
	if pid == nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	project := db.GetProject(tx, pid)
	datasets := db.ListDatasetsFromProject(tx, pid)



	var items []*Dataset

	for _, ds := range datasets{
		items = append(items, &Dataset{
			IsStale:domain.IsStale(ds),
			Item:ds,
		})
	}
	site := shared.LoadSite(tx)
	site.ActiveMenu="projects"
	model := &Model{
		Site:     site,
		Project:  project,
		Datasets: items,
	}

	for _, job := range db.ListProjectJobs(tx, pid){
		job := db.GetJob(tx, job)



		item := JobItem{
			Item: job,

		}


		if job.RunUid != nil {
			run := db.GetJobRun(tx, job.RunUid)
			item.CurrentStatus = run.Status
			item.LastUpdated = run.UpdateTimestamp
		}

		model.Jobs = append(model.Jobs, &item)
	}
	for _, uid := range db.ListProjectSystems(tx, pid){
		model.Systems = append(model.Systems, db.GetSystem(tx, uid))
	}




	h.Layout.Render(w, model)
}
