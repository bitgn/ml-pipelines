package db

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"mlp/catalog/events"
)

func AppendEvent(tx *Tx, e proto.Message) {
	var count uint64

	val := tx.Get(CreateKey(Range_Events))
	if val != nil {
		count = binary.LittleEndian.Uint64(val)

	} else {
		val = make([]byte, 8)
	}

	count +=1
	typ, slice := events.Marshal(e)

	tx.Put(CreateKey(Range_Events, count, int(typ)), slice)

	binary.BigEndian.PutUint64(val, count)

	tx.Put(CreateKey(Range_Events), val)


}