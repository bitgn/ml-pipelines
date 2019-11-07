package view_dashboard

import (
	"mlp/catalog/db"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"

	"net/http"
)



type ProjectModel struct {
	ProjectId string
	Datasets  []*db.DatasetData

}

type Model struct {
	*shared.Site
	Projects []*ProjectModel
	Audit []*AuditModel
	Problems []*AuditModel
	HasMoreAudit bool
	HasMoreProblems bool

}


type AuditModel struct {
	DatasetId,ProjectId string

	Timestamp int64
	Level vo.ACTIVITY_LEVEL
	MultilineText string
}

type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout: tl.DefineTemplate("web/layout.html","web/view_dashboard/content.html"),
	}
}

func (h *Handler) Handle(w http.ResponseWriter){
	tx := h.env.MustRead()

	defer tx.MustAbort()


	var pmod []*ProjectModel

	projects := db.ListProjects(tx)


	for _, p := range projects{

		dss := db.ListDatasets(tx, p.ProjectId)

		pmod = append(pmod, &ProjectModel{
			ProjectId: p.ProjectId,
			Datasets:dss,
		})

	}

	site := shared.LoadSite(tx)
	model := &Model{
		Site: site,
		Projects: pmod,
	}

	const log_limit = 20

	for _, a := range db.ListGlobalActivities(tx, log_limit+1){

		model.Audit = append(model.Audit, &AuditModel{
			MultilineText:a.MultilineText,
			Timestamp:a.UpdateTimestamp,
			Level:a.Level,
			ProjectId:a.ProjectId,
			DatasetId:a.DatasetId,
		})


	}
	if len(model.Audit) > log_limit {
		model.Audit = model.Audit[0:log_limit]
		model.HasMoreAudit = true
	}

	for _, a := range db.ListGlobalProblems(tx, log_limit+1){
		model.Problems = append(model.Problems, &AuditModel{
			MultilineText:a.MultilineText,
			Timestamp:a.UpdateTimestamp,
			Level:a.Level,
			ProjectId:a.ProjectId,
			DatasetId:a.DatasetId,
		})


	}

	if len(model.Problems) > log_limit {
		model.Problems = model.Problems[0:log_limit]
		model.HasMoreProblems = true
	}


	h.layout.Render(w, model)
}