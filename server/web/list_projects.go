package web

import (

	"html/template"
	"mlp/catalog/db"

	"net/http"
)

type ListProjectsModel struct {
	*Site
	Projects []*db.ProjectData
}


func ListProjects(env *db.DB, w http.ResponseWriter){
	tx, err := env.Read()
	if err != nil {
		panic(err)
	}

	defer tx.Close()

	projects := db.ListProjects(tx)



	t, err := template.ParseFiles("web/layout.html","web/list_projects.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	mod := &Site{Url: UrlResolver{}}
	mod.ActiveMenu="projects"

	if err = t.ExecuteTemplate(w, "layout",&ListProjectsModel{mod,projects}); err != nil {
		http.Error(w, err.Error(), 408)
	}
}