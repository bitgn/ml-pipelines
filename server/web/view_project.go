package web

import (
	"html/template"
	"mlp/catalog/db"
	"net/http"
)




type ViewProjectModel struct {
	*Site
	Project *db.ProjectData
	Datasets []*db.DatasetData
}


func ViewProjects(env *db.DB, w http.ResponseWriter, id string){
	tx := env.MustRead()

	defer tx.MustAbort()

	project := db.GetProject(tx, id)
	datasets := db.ListDatasets(tx,project.Id)

	t := template.New("layout")


	t, err := t.ParseFiles("web/layout.html","web/view_projects.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	mod := &Site{}
	mod.ActiveMenu="projects"

	model := &ViewProjectModel{
		Site: mod,
		Project: project,
		Datasets:datasets,
	}
	if err = t.ExecuteTemplate(w, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
		panic(err)
	}
}
