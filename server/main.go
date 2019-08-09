package main

import (
	"errors"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"mlp/catalog/api"
	"mlp/catalog/db"
	"mlp/catalog/sim"
	"mlp/catalog/web"
	"net"
	"net/http"
)



type server struct {
	Env *db.DB
}

func main() {


	cfg := db.NewConfig()
	env, err := db.New("db", cfg)
	if err != nil{
		panic(err)
	}

	defer env.Close()


	s := &server{env}


	mx := mux.NewRouter()
	// static content
	fs := http.FileServer(http.Dir("web/static/"))
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	mx.Path("/explore").Queries("query", "{query:[0-9.\\-A-Za-z]+}").HandlerFunc(simWrap(s.exploreHandler))
	mx.Path("/explore").HandlerFunc(simWrap(s.exploreHandler))
	mx.HandleFunc("/datasets/{dataset_id}/", simWrap(s.datasetHandler))


	mx.Path("/").HandlerFunc(simWrap(s.projectsHandler))



	go runGrpc(env)


	log.Fatal(http.ListenAndServe("localhost:8000", mx))
}

func runGrpc(env *db.DB){
	lis, err := net.Listen("tcp", "localhost:50051")
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



func (s *server) projectsHandler(w http.ResponseWriter, r *http.Request) {
		web.ListProjects(s.Env, w)
}
func (s *server) datasetHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)

	web.ViewDataset(s.Env, w, vars["dataset_id"])
}

func (s *server) exploreHandler(w http.ResponseWriter, r *http.Request){
	query := r.FormValue("query")
	web.ExploreDatasets(s.Env, w, query)
}