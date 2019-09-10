package view_system

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"html/template"
	"mlp/catalog/db"
	"mlp/catalog/graph"
	"mlp/catalog/vo"
	"mlp/catalog/web/shared"
	"net/http"
	"strings"
)

type ViewSystemModel struct {
	*shared.Site
	System     *db.SystemData
	Project     *db.ProjectData
	Lineage     template.HTML
	Description template.HTML
	Kind string


	UsedBy []*SystemLink

}


type SystemLink struct {
	Href string
	Title string
	Entity template.HTML
	Timestamp int64
}

type Handler struct {
	env *db.DB
	layout *shared.LoadedTemplate
}

func NewHandler(env *db.DB, tl *shared.TemplateLoader) *Handler{
	return &Handler{
		env:env,
		layout:tl.DefineTemplate("web/layout.html","web/view_system/content.html"),
	}
}



func (h *Handler) Handle(w http.ResponseWriter, project, service string, version int32){
	tx := h.env.MustRead()

	defer tx.MustAbort()

	pid := db.LookupProject(tx, project)
	if pid == nil {
		http.Error(w, "project not found", http.StatusNotFound)
		return
	}



	ref := db.Lookup(tx, pid, service)
	if ref == nil || ref.Kind != vo.ENTITY_SYSTEM {
		http.Error(w, fmt.Sprintf("service %s not found", service), http.StatusNotFound)
		return
	}




	sys := db.GetSystem(tx, ref.Uid)


	prj := db.GetProject(tx, sys.ProjectUid)


	site := shared.LoadSite(tx)
	site.ActiveMenu="services"


	model := &ViewSystemModel{
		Site:    site,
		System:  sys,
		Project: prj,
		Kind:vo.SystemKind_name[int32(sys.Kind)],
	}


	// lookup users

	links:= db.ListSystemLinks(tx, sys.Uid)

	for _, link := range links{
		switch link.Type {
		case db.SystemLink_Output_SystemVer, db.SystemLink_Input_SystemVer:

			// SVC THIS -->> System deployment
			// service reads from this

			sys := db.GetSystem(tx, link.ContainerUid)
			ver := db.GetSystemVersion(tx, link.InstanceUid)
			model.UsedBy = append(model.UsedBy, &SystemLink{
				Href:      site.Url.ViewSystemVer(sys.ProjectName, sys.Name, ver.VersionNum),
				Title:     sys.Caption(),
				Entity:    site.Fmt.SystemKind(sys.Kind),
				Timestamp: ver.Timestamp,
			})

		case db.SystemLink_Input_JobRun, db.SystemLink_Output_JobRun:
			// SVC THIS ->> JOB run
			user := db.GetJob(tx, link.ContainerUid)
			run := db.GetJobRun(tx, link.InstanceUid)
			model.UsedBy = append(model.UsedBy, &SystemLink{
				Href:site.Url.ViewJobRun(user.ProjectName, user.Name, run.RunNum),
				Title:user.Caption(),
				Entity:"job",
				Timestamp:run.UpdateTimestamp,
			})
		}

	}


	// if version is not provided, we show latest
	if version == 0{
		version = sys.VersionNum
	}

	if version > 0 {
		uid := db.LookupSystemVersion(tx, ref.Uid, version)
		model.Lineage = graph.RenderSystemVersionGraph(tx, site, uid)
	} else {
		model.Lineage= graph.RenderSystemGraph(tx, site, sys.Uid)
	}


	if len(sys.Description) > 0 {
		md := string(markdown.ToHTML([]byte(sys.Description), nil,nil))
		md = strings.Replace(md, "h1", "h4", -1)
		md = strings.Replace(md, "h2", "h5", -1)
		md = strings.Replace(md, "h3", "h6", -1)
		model.Description = template.HTML(md)
	}


	h.layout.Render(w, model)
}

