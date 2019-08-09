package view_dataset

import (
	"bufio"
	"bytes"
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web"
	"net/http"
)

type ViewDatsetModel struct {
	*web.Site
	Dataset     *db.DatasetData
	Project     *db.ProjectData
	IsStale     bool
	Lineage     template.HTML
	Description template.HTML
}




func Handle(env *db.DB, w http.ResponseWriter, datasetId string){
	tx := env.MustRead()

	defer tx.MustAbort()

	ds := db.GetDataset(tx, datasetId)
	pr := db.GetProject(tx, ds.ProjectId)

	t := template.New("layout")


	t, err := t.ParseFiles("web/layout.html","web/view_dataset/content.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	mod := &web.Site{}
	mod.ActiveMenu="projects"

	model := &ViewDatsetModel{
		Site:    mod,
		Dataset: ds,
		Project: pr,
		IsStale: domain.IsStale(ds),
	}

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	if err = t.ExecuteTemplate(foo, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
	} else {
		foo.Flush()
		b.WriteTo(w)

	}
}

