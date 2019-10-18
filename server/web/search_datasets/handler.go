package search_datasets

import (
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web/shared"
	"net/http"
	"strings"
)


type Dataset struct {
	Dataset *db.DatasetData
	IsStale bool

}

type Model struct {
	*shared.Site
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

	ds := meta.Dataset


	if contains(ds.ProjectId, query) {
		return true
	}
	if contains(ds.DatasetId, query){
		return true
	}
	if contains(ds.Summary, query){
		return true
	}
	if contains(ds.Description, query){
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

type Handler struct {
	layout *shared.LoadedTemplate
	env *db.DB
}


func NewHandler(db *db.DB, tl *shared.TemplateLoader) *Handler{


	return &Handler{
		layout:tl.DefineTemplate("web/layout.html","web/search_datasets/content.html"),
		env:db,

	}
}


func (h *Handler) Handle(w http.ResponseWriter, query string){
	tx := h.env.MustRead()

	defer tx.MustAbort()

	datasets := db.ListAllDatasets(tx)

	query = strings.ToLower(strings.TrimSpace(query))

	lines := strings.Split(query, " ")

	var metas []*Dataset

	for _,ds := range datasets{
		meta := &Dataset{
			Dataset:ds,
			IsStale:domain.IsStale(ds),

		}

		if datasetMatchesQuery(meta, lines) {
			metas = append(metas, meta)
		}
	}

	site := shared.LoadSite(tx)
	model := &Model{Site: site, Items: metas, Query: query}

	model.CatalogIsEmpty = len(datasets) == 0

	h.layout.Render(w, model)
}