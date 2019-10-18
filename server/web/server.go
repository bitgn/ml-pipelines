package web

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"mlp/catalog/db"
	"mlp/catalog/web/search_datasets"
	"mlp/catalog/web/view_dashboard"
	"mlp/catalog/web/shared"
	"mlp/catalog/web/view_dataset"
	"net/http"
	"runtime/debug"
)


func NewServer(env *db.DB, templatePath string, devMode bool, specsMode bool, version string) http.Handler{


	tl := shared.NewTemplateLoader(devMode, templatePath)
	shared.SetVersion(version)



	mx := mux.NewRouter()
	// static content
	fs := http.FileServer(http.Dir("web/static/"))
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	{
		handler := search_datasets.NewHandler(env, tl)
		searchFunc := func(w http.ResponseWriter, r *http.Request) {
			query := r.FormValue("query")
			handler.Handle(w, query)

		}
		mx.Path("/search").Queries("query", "{query:[0-9.\\-A-Za-z]+}").HandlerFunc(searchFunc)
		mx.Path("/search").HandlerFunc(searchFunc)

	}

	{
		viewDatasetHandler := view_dataset.NewHandler(env, tl)
		mx.HandleFunc("/projects/{project}/datasets/{dataset}/", func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			viewDatasetHandler.Handle(w, vars["project"], vars["dataset"])
		})
	}

	{
		handler := view_dashboard.NewHandler(env, tl)
		mx.Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler.Handle(w)
		})
	}

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

