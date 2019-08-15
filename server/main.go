package main

import (
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"mlp/catalog/api"
	"mlp/catalog/core"
	"mlp/catalog/db"
	"mlp/catalog/sim"
	"mlp/catalog/web"
	"mlp/catalog/web/explore_datasets"
	"mlp/catalog/web/list_projects"
	"mlp/catalog/web/view_dataset"
	"mlp/catalog/web/view_project"
	"net"
	"net/http"
)



type server struct {
	Env *db.DB
	version string
}

var (
	webInterface = flag.String("web", "localhost:8080", "web interface to bind to")
	grpcInterface = flag.String("grpc", "localhost:9111", "GRPC interface to bind to")
	dbFolder = flag.String("db", "db", "Folder to store local database")
	devMode = flag.Bool("dev", false, "Enable dynamic template reloading")
	testMode = flag.Bool("test", false, "Enable test server and use async LMDB mode")
	upgrade = flag.String("upgrade", "auto", "Upgrade projections: auto, force, none")
	version = "dev"
)


func main() {


	log.Printf("Starting MLP-Catalog %s", version)

	web.SetVersion(version)
	flag.Parse()




	if *devMode {
		log.Println("Enable template reloading")
		web.AlwaysReloadTemplates()
	}

	cfg := db.NewConfig()

	cfg.TestMode = *testMode
	env, err := db.New(*dbFolder, cfg)
	if err != nil {
		panic(err)
	}

	defer env.Close()


	switch *upgrade {
	case "auto":
		core.UpgradeDB(env, version, core.UpgradePolicy_Auto)
	case "force":
		core.UpgradeDB(env, version, core.UpgradePolicy_Force)
	}


	s := &server{Env: env, version: version}

	mx := mux.NewRouter()
	// static content
	fs := http.FileServer(http.Dir("web/static/"))
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// explore datasets
	mx.Path("/explore").Queries("query", "{query:[0-9.\\-A-Za-z]+}").HandlerFunc(simWrap(s.exploreHandler))
	mx.Path("/explore").HandlerFunc(simWrap(s.exploreHandler))
	// view dataset
	mx.HandleFunc("/datasets/{dataset_id}", simWrap(s.datasetHandler))
	// vie project
	mx.HandleFunc("/projects/{project_id}", simWrap(s.projectHandler))

	mx.Path("/").HandlerFunc(simWrap(s.projectsHandler))



	log.Printf("Starting web at %s gRPC at %s\n", *webInterface, *grpcInterface)

	go runGrpc(env)

	log.Fatal(http.ListenAndServe(*webInterface, mx))
}


func runGrpc(env *db.DB){
	lis, err := net.Listen("tcp", *grpcInterface)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}




	grpcServer := grpc.NewServer()

	//if *testMode {
	// TEMP for now
		testService := api.NewTestServer(env)
		api.RegisterTestServer(grpcServer, testService)
	//}

	catalogService := api.NewCatalogServer(env)
	api.RegisterCatalogServer(grpcServer, catalogService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func simWrap(inner http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		if sim.IsRunning(){

			var err error
			defer func() {
				rec := recover()
				if rec != nil {
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
			//fmt.Println("  ", r.URL)
		}
		inner(w, r)
	}
}

func (s *server) projectHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	view_project.Handle(s.Env, w, vars["project_id"])
}

func (s *server) projectsHandler(w http.ResponseWriter, r *http.Request) {
		list_projects.Handle(s.Env, w)
}
func (s *server) datasetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	view_dataset.Handle(s.Env, w, vars["dataset_id"])
}

func (s *server) exploreHandler(w http.ResponseWriter, r *http.Request){
	query := r.FormValue("query")
	explore_datasets.Handle(s.Env, w, query)
}