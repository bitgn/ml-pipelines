package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/mlp_api"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
	"mlp/catalog/test_api"
	"mlp/catalog/web"
	"net"
	"net/http"
	"os"
)





var (

	fs = flag.NewFlagSet("mlp-server", flag.ExitOnError)

	webInterface  = fs.String("web", "localhost:8080", "web interface to bind to")
	grpcInterface = fs.String("grpc", "localhost:9111", "GRPC interface to bind to")
	dbFolder      = fs.String("db", "db", "Folder to store local database")
	devMode       = fs.Bool("dev", false, "Dev features: template reloading, asserts")
	testAPI       = fs.Bool("test-api", false, "Enable test server")
	upgrade       = fs.String("upgrade", "auto", "Upgrade projections: auto, force, none")
	specsMode     = fs.Bool("specs", false, "Enable spec-runner mode")
	templatePath  = fs.String("template-path", "", "HTML Layout root")

	// build script will replace this one
	version = "dev"
)


func main() {


	//log.Printf("Starting MLP-Catalog %s with args %v/n", version, os.Args)



	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatalln("Error", err.Error())
	}

	if *specsMode{
		// test API is implied for specs
		*testAPI = true
	}


	cfg := db.NewConfig()

	cfg.Async = *specsMode
	env, err := db.New(*dbFolder, cfg)
	if err != nil {
		panic(err)
	}

	defer env.Close()


	switch *upgrade {
	case "auto":
		projection.UpgradeDB(env, version, projection.UpgradePolicy_Auto)
	case "force":
		projection.UpgradeDB(env, version, projection.UpgradePolicy_Force)
	}



	log.Printf("Starting web at %s gRPC at %s\n", *webInterface, *grpcInterface)

	go runGrpc(env)

	if *specsMode{
		sim.Start()
	}

	server := web.NewServer(env, *templatePath, *devMode, *specsMode, version)

	log.Fatal(http.ListenAndServe(*webInterface, server))
}


func runGrpc(env *db.DB){
	lis, err := net.Listen("tcp", *grpcInterface)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	if *testAPI {
		log.Println("Enabling TEST interface")
		testService := test_api.NewServer(env)
		test_api.RegisterTestServer(grpcServer, testService)
	}

	catalogService := mlp_api.NewServer(env, version)
	mlp_api.RegisterCatalogServer(grpcServer, catalogService)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
