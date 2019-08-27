package sim

import (
	"encoding/binary"
	"github.com/google/uuid"
)

func NewID() []byte {

	if sim_running{



		buf := make([]byte, 16)
		binary.BigEndian.PutUint64(buf[8:],sim_counter)

		sim_counter+=1
		return buf


	} else {
		id := uuid.Must(uuid.NewUUID())
		return id[:]
	}



}
