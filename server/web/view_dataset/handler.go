package view_dataset

import (
	"bufio"
	"bytes"
	"github.com/gomarkdown/markdown"
	"html/template"
	"io"
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
	Lineage     template.HTML
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



func (h *Handler) Handle(w http.ResponseWriter, datasetId string){
	tx := h.env.MustRead()

	defer tx.MustAbort()

	ds := db.GetDataset(tx, datasetId)
	pr := db.GetProject(tx, ds.ProjectId)


	site := shared.LoadSite(tx)
	site.ActiveMenu="projects"

	model := &ViewDatsetModel{
		Site:    site,
		Dataset: ds,
		Project: pr,
		IsStale: domain.IsStale(ds),
		Lineage: renderSVG( tx,datasetId, site.Url),
	}

	if len(ds.Description) > 0 {
		md := string(markdown.ToHTML([]byte(ds.Description), nil,nil))
		md = strings.Replace(md, "h1", "h4", -1)
		md = strings.Replace(md, "h2", "h5", -1)
		md = strings.Replace(md, "h3", "h6", -1)
		model.Description = template.HTML(md)
	}


	io.Pipe()

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	if err := h.layout.Exec(foo, model); err != nil {
		http.Error(w, err.Error(), 408)
	} else {
		foo.Flush()
		b.WriteTo(w)

	}
}

