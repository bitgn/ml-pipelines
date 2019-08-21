package list_projects

import (
	"mlp/catalog/db"
	"mlp/catalog/web/shared"

	"net/http"
)

type Model struct {
	*shared.Site
	Projects []*db.ProjectData
}


type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout: tl.DefineTemplate("web/layout.html","web/list_projects/content.html"),
	}
}

func (h *Handler) Handle(w http.ResponseWriter){
	tx := h.env.MustRead()

	defer tx.MustAbort()

	projects := db.ListProjects(tx)


	site := shared.LoadSite(tx)
	site.ActiveMenu = "projects"

	model := &Model{
		Site: site,
		Projects: projects,
	}

	if err:= h.layout.Exec(w, model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}