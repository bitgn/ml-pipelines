package explore_datasets

import (
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web"
	"net/http"
	"strings"
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



func contains(source, lowercaseSubstring string) bool{
	return strings.Contains(strings.ToLower(source), lowercaseSubstring)
}

func datasetMatchesOne(meta *Dataset, query string) bool {
	if query == "stale"  && meta.IsStale{
			return true
	}

	if contains(meta.ProjectName, query) {
		return true
	}
	ds := meta.Dataset
	if contains(ds.Name, query){
		return true
	}
	if contains(ds.DataFormat, query){
		return true
	}
	if contains(ds.Description, query){
		return true
	}
	if contains(ds.LocationId, query){
		return true
	}
	if contains(ds.LocationUri, query){
		return true
	}
	if ds.Sample != nil && contains(string(ds.Sample.Body), query){
		return true
	}

	return false
}

func datasetMatchesQuery(ds *Dataset, query []string) bool {

	if len(query) == 0{
		return true
	}

	for _, l := range query {
		if !datasetMatchesOne(ds, l){
			return false
		}
	}
	return true
}


var layout = web.DefineTemplate("web/layout.html","web/explore_datasets/content.html")

func Handle(env *db.DB, w http.ResponseWriter, query string){
	tx := env.MustRead()

	defer tx.MustAbort()

	datasets := db.ListAllDatasets(tx)

	query = strings.ToLower(strings.TrimSpace(query))

	lines := strings.Split(query, " ")

	var metas []*Dataset

	for _,ds := range datasets{
		prj := db.GetProject(tx, ds.ProjectId)
		meta := &Dataset{
			Dataset:ds,
			IsStale:domain.IsStale(ds),
			ProjectName:prj.Name,
		}

		if datasetMatchesQuery(meta, lines) {
			metas = append(metas, meta)
		}
	}

	site := web.LoadSite(tx)
	site.ActiveMenu="datasets"

	model := &Model{Site: site, Items: metas, Query: query}

	model.CatalogIsEmpty = len(datasets) == 0

	if err := layout.ExecuteTemplate(w, model); err != nil {
		http.Error(w, err.Error(), 408)
	}
}