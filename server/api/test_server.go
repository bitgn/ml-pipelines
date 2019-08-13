package api

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
	"os"
)

//Setup(context.Context, *ScenarioRequest) (*ScenarioResponse, error)


type test_server struct {

	// TODO: async DB would be faster for tests
	db *db.DB

}

func (t *test_server) Wipe(context.Context, *WipeDatabase) (*OkResponse, error) {
	tx := t.db.MustWrite()

	defer tx.MustCleanup()

	if err := tx.Tx.Drop(tx.DB, false); err != nil {
		panic(errors.Wrap(err, "Failed to empty the DB"))
	}
	tx.MustCommit()
	return &OkResponse{}, nil
}

func (t *test_server) Ping(context.Context, *PingRequest) (*OkResponse, error) {
	return &OkResponse{}, nil
}

func (t *test_server) Setup(ctx context.Context, req *ScenarioRequest) (*OkResponse, error){
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

func (t *test_server) Kill(ctx context.Context, req *KillRequest) (*OkResponse, error){
	os.Exit(0)
	return &OkResponse{}, nil
}


func NewTestServer(db *db.DB) TestServer{
	return &test_server{db}
}



