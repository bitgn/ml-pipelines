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

	for _, a := range db.ListGlobalActivities(tx){



		model.Audit = append(model.Audit, &AuditModel{
			MultilineText:a.MultilineText,
			Timestamp:a.UpdateTimestamp,
			Level:a.Level,
			ProjectId:a.ProjectId,
			DatasetId:a.DatasetId,

		})
	}

	h.layout.Render(w, model)
}