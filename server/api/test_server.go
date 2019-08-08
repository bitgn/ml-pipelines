package api

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/projection"
	"mlp/catalog/sim"
)

//Setup(context.Context, *ScenarioRequest) (*ScenarioResponse, error)


type test_server struct {

	// TODO: async DB would be faster for tests
	db *db.DB

}




func (t *test_server) Setup(ctx context.Context, req *ScenarioRequest) (*ScenarioResponse, error){
	fmt.Printf("Received scenario '%s' with %d events\n", req.Name, len(req.Events))

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
	}

	tx.MustCommit()

	return &ScenarioResponse{}, nil
}


func NewTestServer(db *db.DB) TestServer{
	return &test_server{db}

}

