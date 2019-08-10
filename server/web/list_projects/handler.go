package list_projects

import (
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/web"

	"net/http"
)

type Model struct {
	*web.Site
	Projects []*db.ProjectData
}


func Handle(env *db.DB, w http.ResponseWriter){
	tx := env.MustRead()

	defer tx.MustAbort()

	projects := db.ListProjects(tx)

	t := template.New("layout")


	t, err := t.ParseFiles("web/layout.html","web/list_projects/content.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	site := web.LoadSite(tx)
	site.ActiveMenu = "projects"

	model := &Model{
		Site: site,
		Projects: projects,
	}
	if err = t.ExecuteTemplate(w, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}