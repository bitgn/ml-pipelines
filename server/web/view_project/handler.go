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


func (h *Handler) Handle(w http.ResponseWriter, id string){
	tx := h.Env.MustRead()

	defer tx.MustAbort()

	project := db.GetProject(tx, id)
	datasets := db.ListDatasets(tx,project.Id)



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
	var err error
	if err = h.Layout.Exec(w, model); err != nil {
		http.Error(w, err.Error(), 408)
		panic(err)
	}
}
