package db

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"mlp/catalog/events"
)

func AppendEvent(tx *Tx, e proto.Message) uint64 {
	var version uint64

	val := tx.Get(CreateKey(Range_Events))
	if val != nil {
		version = binary.LittleEndian.Uint64(val)

	} else {
		val = make([]byte, 8)
	}

	version +=1
	typ, slice := events.Marshal(e)

	tx.Put(CreateKey(Range_Events, version, int(typ)), slice)

	binary.BigEndian.PutUint64(val, version)

	tx.Put(CreateKey(Range_Events), val)

	return version


}