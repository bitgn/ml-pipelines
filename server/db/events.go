package db

import (
	"encoding/binary"
	"github.com/abdullin/lex-go/tuple"
	"github.com/golang/protobuf/proto"
	"log"
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




type Projector func(*Tx, proto.Message);

func ReplayEvents(tx *Tx, p Projector) int{
	count := 0
	updater := func (k, v []byte){

		t, err := tuple.Unpack(k)

		if err != nil {
			log.Panicln("Failed to unpack tuple", err)
		}

		if len(t) != 3 {
			return
		}

		kind, ok := t[2].(int64)
		if !ok {
			log.Panicln("Failed to parse event type from the db:", t[2])
		}

		i := events.Type(kind)
		event := events.Unmarshal(i, v)

		defer func() {
			if r := recover(); r != nil {
				log.Panicln("Failure while processing ", events.Type_name[int32(i)], "body", event, r)
			}
		}()

		p(tx, event)

		count +=1

	}

	tx.MustScanRange(CreateKey(Range_Events),updater)

	return count


}



