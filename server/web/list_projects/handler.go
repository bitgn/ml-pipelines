package list_projects

import (
	"mlp/catalog/db"
	"mlp/catalog/web"

	"net/http"
)

type Model struct {
	*web.Site
	Projects []*db.ProjectData
}

var layout = web.DefineTemplate("web/layout.html","web/list_projects/content.html")


func Handle(env *db.DB, w http.ResponseWriter){
	tx := env.MustRead()

	defer tx.MustAbort()

	projects := db.ListProjects(tx)


	site := web.LoadSite(tx)
	site.ActiveMenu = "projects"

	model := &Model{
		Site: site,
		Projects: projects,
	}

	if err:= layout.ExecuteTemplate(w, model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}