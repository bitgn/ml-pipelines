package view_job

import (
	"fmt"
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
	"net/http"
)

type ViewJobModel struct {
	*shared.Site
	Item        *db.Job
	Project     *db.ProjectData
	Runs []*db.JobRunData
	Lineage     template.HTML
	Description template.HTML


	CurrentStatus vo.JOB_STATUS
	LastUpdated int64

	LastSuccessTimestamp string

}
type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout:tl.DefineTemplate("web/layout.html","web/view_job/content.html"),
	}
}



func (h *Handler) Handle(w http.ResponseWriter, project_name, job_name string){
	tx := h.env.MustRead()

	defer tx.MustAbort()

	pid := db.LookupProject(tx, project_name)
	if pid == nil {
		http.Error(w, fmt.Sprintf("project %s not found", project_name), http.StatusNotFound)
		return
	}



	ref := db.Lookup(tx, pid, job_name)
	if ref == nil || ref.Kind != vo.ENTITY_JOB {
		http.Error(w, fmt.Sprintf("job %s not found", job_name), http.StatusNotFound)
		return
	}




	job := db.GetJob(tx, ref.Uid)
	prj := db.GetProject(tx, job.ProjectUid)


	site := shared.LoadSite(tx)
	site.ActiveMenu="jobs"


	model := &ViewJobModel{
		Site:    site,
		Item:  job,
		Project: prj,
	}

	if job.RunUid != nil {
		lastRun := db.GetJobRun(tx, job.RunUid)
		model.CurrentStatus = lastRun.Status
		model.LastUpdated = lastRun.UpdateTimestamp
	}

	if job.LastSuccessUid != nil {
		lastSuccess := db.GetJobRun(tx, job.LastSuccessUid)
		model.LastSuccessTimestamp = site.Fmt.Timestamp(lastSuccess.UpdateTimestamp)
	} else {
		model.LastSuccessTimestamp = "never"
	}




	var runs  []*db.JobRunData

	for i := job.RunNum; i>0; i-- {
		runs = append(runs, db.LookupJobRun(tx, job.Uid, i))
		if len(runs)>=10{
			break
		}

	}
	model.Runs = runs




	h.layout.Render(w, model)
}

