package view_project

import (
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web"
	"net/http"
)


type Dataset struct {
	Item *db.DatasetData
	IsStale bool
}


type Model struct {
	*web.Site
	Project *db.ProjectData
	Datasets []*Dataset
}


func Handle(env *db.DB, w http.ResponseWriter, id string){
	tx := env.MustRead()

	defer tx.MustAbort()

	project := db.GetProject(tx, id)
	datasets := db.ListDatasets(tx,project.Id)

	t := template.New("layout")

	t, err := t.ParseFiles("web/layout.html","web/view_project/content.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	var items []*Dataset

	for _, ds := range datasets{
		items = append(items, &Dataset{
			IsStale:domain.IsStale(ds),
			Item:ds,
		})
	}

	mod := web.LoadSite(tx)
	mod.ActiveMenu="projects"

	model := &Model{
		Site: mod,
		Project: project,
		Datasets:items,
	}
	if err = t.ExecuteTemplate(w, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
		panic(err)
	}
}
