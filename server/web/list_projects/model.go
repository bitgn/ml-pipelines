package list_projects

import (

	"html/template"
	"mlp/catalog/db"

	"net/http"
)

type Model struct {
	Projects []*db.ProjectData

}




func Handle(env *db.DB, w http.ResponseWriter){
	tx, err := env.Read()
	if err != nil {
		panic(err)
	}

	defer tx.Close()

	projects := db.ListProjects(tx)

	t, err := template.ParseFiles("web/list_projects/view.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}
	t.Execute(w, &Model{projects})
}