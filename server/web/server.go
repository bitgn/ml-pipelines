package web

import (
	"encoding/hex"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"mlp/catalog/db"
	"mlp/catalog/web/explore_datasets"
	"mlp/catalog/web/list_projects"
	"mlp/catalog/web/shared"
	"mlp/catalog/web/view_dataset"
	"mlp/catalog/web/view_project"
	"mlp/catalog/web/view_service"
	"net/http"
	"runtime/debug"
	"strconv"
)


func NewServer(env *db.DB, templatePath string, devMode bool, specsMode bool, version string) http.Handler{


	tl := shared.NewTemplateLoader(devMode, templatePath)
	shared.SetVersion(version)



	mx := mux.NewRouter()
	// static content
	fs := http.FileServer(http.Dir("web/static/"))
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))



	exploreHandler := explore_datasets.NewHandler(env, tl)
	exploreFunc := func(w http.ResponseWriter, r *http.Request){
		query := r.FormValue("query")
		exploreHandler.Handle(w, query)

	}
	mx.Path("/explore").Queries("query", "{query:[0-9.\\-A-Za-z]+}").HandlerFunc(exploreFunc)
	mx.Path("/explore").HandlerFunc(exploreFunc)



	projectHandler := view_project.NewHandler(env, tl)
	mx.HandleFunc("/projects/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectHandler.Handle(w, vars["name"])
	})



	viewDatasetHandler := view_dataset.NewHandler(env, tl)
	mx.HandleFunc("/projects/{project}/datasets/{dataset}/ver/{version}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ver, err := hex.DecodeString(vars["version"])
		if err != nil {
			http.Error(w, "Dataset version is in invalid format", http.StatusBadRequest)
			return
		}
		viewDatasetHandler.Handle(w, vars["project"], vars["dataset"], ver)
	})
	mx.HandleFunc("/projects/{project}/datasets/{dataset}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		viewDatasetHandler.Handle(w, vars["project"], vars["dataset"], nil)
	})


	viewServiceHandler := view_service.NewHandler(env, tl)
	mx.HandleFunc("/projects/{project}/services/{service}/ver/{version}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		ver, err := strconv.Atoi(vars["version"])
		if err != nil {
			http.Error(w, "Version is in invalid format", http.StatusBadRequest)
			return
		}
		viewServiceHandler.Handle(w, vars["project"], vars["service"], int32(ver))
	})
	mx.HandleFunc("/projects/{project}/services/{service}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		viewServiceHandler.Handle(w, vars["project"], vars["service"], 0)
	})



	listProjectsHandler := list_projects.NewHandler(env, tl)
	mx.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		listProjectsHandler.Handle(w)
	})

	if specsMode{
		return wrapError(mx.ServeHTTP)
	}
	return http.HandlerFunc(mx.ServeHTTP)
}



func wrapError(inner http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		var err error
		defer func() {
			rec := recover()
			if rec != nil {
				println(string(debug.Stack()))
				if r != nil {
					switch t := rec.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("Unknown error")
					}
					println(err.Error())

					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}

		}()
		inner(w, r)
	}
}

