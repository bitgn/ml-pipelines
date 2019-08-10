package main

import (
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"mlp/catalog/api"
	"mlp/catalog/db"
	"mlp/catalog/sim"
	"mlp/catalog/web/explore_datasets"
	"mlp/catalog/web/list_projects"
	"mlp/catalog/web/view_dataset"
	"mlp/catalog/web/view_project"
	"net"
	"net/http"
)



type server struct {
	Env *db.DB
}

var (
	webInterface = flag.String("web", "localhost:8080", "web interface to bind to")
	grpcInterface = flag.String("grpc", "localhost:50051", "GRPC interface to bind to")
	dbFolder = flag.String("db", "db", "Folder to store local database")
)


func main() {


	flag.Parse()


	cfg := db.NewConfig()
	env, err := db.New(*dbFolder, cfg)
	if err != nil{
		panic(err)
	}

	defer env.Close()


	s := &server{env}


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



	go runGrpc(env)


	log.Fatal(http.ListenAndServe(*webInterface, mx))
}

func runGrpc(env *db.DB){
	lis, err := net.Listen("tcp", *grpcInterface)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	testService := api.NewTestServer(env)

	grpcServer := grpc.NewServer()

	api.RegisterTestServer(grpcServer, testService)
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
						print(err.Error())
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