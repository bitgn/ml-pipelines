package view_job_run

import (
	"fmt"
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/graph"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
	"net/http"
)

type Model struct {
	*shared.Site
	Run     *db.JobRunData
	Job     *db.Job
	Project *db.ProjectData
	Lineage template.HTML
	History []*Event

}


type Event struct {
	Class string
	Timestamp int64
	Text string
}
type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout:tl.DefineTemplate("web/layout.html","web/view_job_run/content.html"),
	}
}



func (h *Handler) Handle(w http.ResponseWriter, project_name, job_name string, run_num int32){
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
	run := db.LookupJobRun(tx, ref.Uid, run_num)


	site := shared.LoadSite(tx)
	site.ActiveMenu="jobs"


	model := &Model{
		Site:    site,
		Job:     job,
		Project: prj,
		Run:     run,
	}

	model.Lineage =graph.RenderJobRunGraph(tx, site, run.Uid)
	for _, e := range db.ListJobRunHistory(tx, run.Uid){

		model.History = append(model.History, toModel(e))

	}



	h.layout.Render(w, model)
}


func toModel(e *db.JobRunEvent) *Event{
	switch e.Type {
	case db.JobRunEvent_Started:
		return &Event{
			Class:"info",
			Timestamp:e.Timestamp,
			Text:"Started",

		}
	case db.JobRunEvent_Failed:
		return &Event{
			Class:"danger",
			Timestamp:e.Timestamp,
			Text:e.Log,
		}

	case db.JobRunEvent_Logged:
		return &Event{
			Class:"fail",
			Timestamp:e.Timestamp,
			Text:e.Log,
		}
	case db.JobRunEvent_Completed:
		return &Event{
			Class:"success",
			Timestamp:e.Timestamp,
			Text:"Completed",
		}
	default:
		panic(e.Type)
	}

}

