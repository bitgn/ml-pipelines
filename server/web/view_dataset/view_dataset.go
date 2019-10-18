package view_dataset

import (
	"github.com/gomarkdown/markdown"
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/web/shared"
	"net/http"
	"strings"
)

type ViewDatsetModel struct {
	*shared.Site
	Dataset     *db.DatasetData
	Project     *db.ProjectData
	IsStale     bool
	Description template.HTML
}

type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout:tl.DefineTemplate("web/layout.html","web/view_dataset/content.html"),
	}
}



func (h *Handler) Handle(w http.ResponseWriter, project, dataset string){
	tx := h.env.MustRead()

	defer tx.MustAbort()


	prj := db.GetProject(tx, project)

	if prj == nil {
		http.Error(w, "project not found",http.StatusNotFound)
		return
	}

	ds := db.GetDataset(tx, project, dataset)
	if ds == nil {
		http.Error(w, "dataset not found", http.StatusNotFound)
		return
	}

	site := shared.LoadSite(tx)

	model := &ViewDatsetModel{
		Site:    site,
		Dataset: ds,
		Project: prj,
		IsStale: domain.IsStale(ds),
	}

	if len(ds.Description) > 0 {
		md := string(markdown.ToHTML([]byte(ds.Description), nil,nil))
		md = strings.Replace(md, "h1", "h4", -1)
		md = strings.Replace(md, "h2", "h5", -1)
		md = strings.Replace(md, "h3", "h6", -1)
		model.Description = template.HTML(md)
	}


	h.layout.Render(w, model)
}

