package web

import (
	"fmt"
	"html/template"
	"mlp/catalog/db"

	"net/http"
)

type ListProjectsModel struct {
	*Site
	Projects []*db.ProjectData
}


func ListProjects(env *db.DB, w http.ResponseWriter){
	tx := env.MustRead()

	defer tx.MustAbort()

	projects := db.ListProjects(tx)

	fmt.Println("Downloaded %d", len(projects))



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