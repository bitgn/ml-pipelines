package view_dataset

import (
	"github.com/gomarkdown/markdown"
	"html/template"
	"io"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/vo"
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



func (h *Handler) Handle(w http.ResponseWriter, project, dataset string, verUid []byte){
	tx := h.env.MustRead()

	defer tx.MustAbort()



	did := db.Lookup(tx, project, dataset)
	if did == nil || did.Kind != vo.ENTITY_DATASET {
		http.Error(w, "dataset not found", http.StatusNotFound)
		return
	}




	ds := db.GetDataset(tx, did.Uid)
	pr := db.GetProject(tx, ds.ProjectUid)


	site := shared.LoadSite(tx)
	site.ActiveMenu="projects"


	var ver *db.DatasetVersionData
	if verUid == nil {
		ver = db.GetLastDatasetVersion(tx, did.Uid)
		if ver != nil {
			verUid = ver.Uid
		}
	} else {
		ver = db.GetDatasetVersion(tx, verUid)
	}




	model := &ViewDatsetModel{
		Site:    site,
		Dataset: ds,
		Project: pr,
		IsStale: domain.IsStale(ds),
		Lineage: renderDatasetVersionSVG( tx,verUid, site.Url),
	}

	if len(ds.Description) > 0 {
		md := string(markdown.ToHTML([]byte(ds.Description), nil,nil))
		md = strings.Replace(md, "h1", "h4", -1)
		md = strings.Replace(md, "h2", "h5", -1)
		md = strings.Replace(md, "h3", "h6", -1)
		model.Description = template.HTML(md)
	}


	io.Pipe()

	h.layout.Render(w, model)
}

