package web

import (
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"net/http"
)


type ExploreDatasetItem struct {
	Dataset *db.DatasetData
	IsStale bool
	ProjectName string
}

type ExploreDatasetsModel struct {
	*Site
	Items []*ExploreDatasetItem
	Query string
	CatalogIsEmpty bool
}


func ExploreDatasets(env *db.DB, w http.ResponseWriter, query string){
	tx := env.MustRead()

	defer tx.MustAbort()

	datasets := db.ListAllDatasets(tx)

	t := template.New("layout")



	var metas []*ExploreDatasetItem

	for _,ds := range datasets{
		meta := &ExploreDatasetItem{
			Dataset:ds,
			IsStale:domain.IsStale(ds),
		}

		metas = append(metas, meta)
	}


	t, err := t.ParseFiles("web/layout.html","web/explore_datasets.html")
	if err != nil{
		http.Error(w, err.Error(), 408)
		return
	}

	mod := &Site{}
	mod.ActiveMenu="datasets"

	model := &ExploreDatasetsModel{Site: mod, Items: metas, Query: query}

	model.CatalogIsEmpty = len(datasets) == 0



	if err = t.ExecuteTemplate(w, "layout", model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}

