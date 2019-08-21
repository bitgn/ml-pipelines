package test_api

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
	"os"
)

//Setup(context.Context, *ScenarioRequest) (*ScenarioResponse, error)


type server struct {
	db *db.DB
}

func NewServer(db *db.DB) TestServer{
	return &server{db}
}



func (c *server) Apply(ctx context.Context, req *ApplyRequest) (*ApplyResponse, error) {
	tx := c.db.MustWrite()
	defer tx.MustCleanup()

	var version uint64

	for _, e := range req.Events{
		msg := events.Unmarshal(e.Type, e.Body)
		version = c.publish(tx, msg)

	}
	tx.MustCommit()

	return &ApplyResponse{
		Version:version,
	}, nil
}


func (c *server) publish(tx *db.Tx, e proto.Message) uint64{
	projection.Handle(tx, e)
	return db.AppendEvent(tx, e)
}


func (t *server) Wipe(context.Context, *WipeDatabase) (*OkResponse, error) {

	print("Open tx to delete")
	tx := t.db.MustWrite()

	defer tx.MustCleanup()

	print("Ready to delete")

	if err := tx.Tx.Drop(tx.DB, false); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to empty the DB"))
	}
	tx.MustCommit()
	return &OkResponse{}, nil
}

func (t *server) Ping(context.Context, *PingRequest) (*OkResponse, error) {
	return &OkResponse{}, nil
}

func (t *server) Setup(ctx context.Context, req *ScenarioRequest) (*OkResponse, error){
	//fmt.Printf("Received scenario '%s' with %d events\n", req.Name, len(req.Events))

	sim.Start()

	tx := t.db.MustWrite()

	if err := tx.Tx.Drop(tx.DB, false); err != nil {
		panic(errors.Wrap(err, "Failed to empty the DB"))
	}

	if req.Timestamp != 0 {
		sim.SetSimulationTime(req.Timestamp)
	} else {
		sim.DisableSimulationTime()
	}


	for _,e := range req.Events{
		msg := events.Unmarshal(e.Type, e.Body)
		projection.Handle(tx, msg)
		db.AppendEvent(tx, msg)
	}

	tx.MustCommit()

	return &OkResponse{}, nil
}

func (t *server) Kill(ctx context.Context, req *KillRequest) (*OkResponse, error){
	os.Exit(0)
	return &OkResponse{}, nil
}






