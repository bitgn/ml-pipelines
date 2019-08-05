package main

import (
	"log"
	"mlp/catalog/db"
	"mlp/catalog/web"

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



	http.HandleFunc("/", s.projectsHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}




func (s *server) projectsHandler(w http.ResponseWriter, r *http.Request) {
		web.ListProjects(s.Env, w)

}