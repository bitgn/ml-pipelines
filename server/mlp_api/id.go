package mlp_api

import "github.com/google/uuid"

func newID() []byte {
	id := uuid.Must(uuid.NewUUID())
	return id[:]

}
