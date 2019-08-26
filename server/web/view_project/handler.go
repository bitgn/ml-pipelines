package view_project

import (
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web/shared"
	"net/http"
)


type Dataset struct {
	Item *db.DatasetData
	IsStale bool
}


type Model struct {
	*shared.Site
	Project *db.ProjectData
	Datasets []*Dataset
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

	mod := shared.LoadSite(tx)
	mod.ActiveMenu="projects"

	model := &Model{
		Site: mod,
		Project: project,
		Datasets:items,
	}

	h.Layout.Render(w, model)
}
