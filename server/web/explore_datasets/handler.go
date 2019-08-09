package explore_datasets

import (
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web"
	"net/http"
)


type Dataset struct {
	Dataset *db.DatasetData
	IsStale bool
	ProjectName string
}

type Model struct {
	*web.Site
	Items []*Dataset
	Query string
	CatalogIsEmpty bool
}


func Handle(env *db.DB, w http.ResponseWriter, query string){
	tx := env.MustRead()

	defer tx.MustAbort()

	datasets := db.ListAllDatasets(tx)

	t := template.New("layout")



	var metas []*Dataset

	for _,ds := range datasets{
		meta := &Dataset{
			Dataset:ds,
			IsStale:domain.IsStale(ds),
		}

		metas = append(metas, meta)
	}


	t, err := t.ParseFiles("web/layout.html","web/explore_datasets/content.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	mod := &web.Site{}
	mod.ActiveMenu="datasets"

	model := &Model{Site: mod, Items: metas, Query: query}

	model.CatalogIsEmpty = len(datasets) == 0



	if err = t.ExecuteTemplate(w, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}

