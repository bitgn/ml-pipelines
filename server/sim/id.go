package sim

import (
	"encoding/binary"
	"github.com/google/uuid"
)

func NewID() []byte {

	if sim_running{


		sim_counter+=1
		buf := make([]byte, 16)
		binary.BigEndian.PutUint64(buf[8:],sim_counter)


		return buf


	} else {
		id := uuid.Must(uuid.NewUUID())
		return id[:]
	}



}
