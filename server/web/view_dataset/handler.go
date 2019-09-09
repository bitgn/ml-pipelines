package view_dataset

import (
	"encoding/hex"
	"fmt"
	"github.com/gomarkdown/markdown"
	"html/template"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/domain"
	"mlp/catalog/graph"
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

	pid := db.LookupProject(tx, project)
	if pid == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}



	did := db.Lookup(tx, pid, dataset)
	if did == nil || did.Kind != vo.ENTITY_DATASET {
		http.Error(w, "dataset not found", http.StatusNotFound)
		return
	}




	ds := db.GetDataset(tx, did.Uid)
	pr := db.GetProject(tx, ds.ProjectUid)


	site := shared.LoadSite(tx)
	site.ActiveMenu="projects"


	var ver *db.DatasetVersionData
	if verUid != nil {
		ver = db.GetDatasetVersion(tx, verUid)
		if ver == nil {
			http.Error(w, fmt.Sprintf("Unknown version %s", hex.EncodeToString(verUid)), http.StatusBadRequest)
		}
	} else {

		if ds.VersionCount > 0 {
			ver = db.GetDatasetVersion(tx, ds.HeadVersion)
			if ver == nil {
				log.Panicln("Can't load HEAD version that should exist")
			}

				verUid = ver.Uid
		}

	}


	render := graph.RenderDatasetVerGraph(tx, site, verUid)


	model := &ViewDatsetModel{
		Site:    site,
		Dataset: ds,
		Project: pr,
		IsStale: domain.IsStale(ds),
		Lineage: render,
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

