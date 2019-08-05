package main

import (
	"google.golang.org/grpc"
	"log"
	"mlp/catalog/api"
	"mlp/catalog/db"
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


	fs := http.FileServer(http.Dir("web/static/"))


	http.Handle("/static/", http.StripPrefix("/static/", fs))


	go runGrpc(env)

	http.HandleFunc("/", s.projectsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func runGrpc(env *db.DB){
	lis, err := net.Listen("tcp", ":8001")
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




func (s *server) projectsHandler(w http.ResponseWriter, r *http.Request) {
		web.ListProjects(s.Env, w)

}