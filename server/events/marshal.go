package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)





func getInstance(id Event_Types) proto.Message{
	switch id {
	case Event_ProjectCreated:
		return &ProjectCreated{}
	case Event_DatasetCreated:
		return &DatasetCreated{}

	default:
		panic(errors.Errorf("Unknown event type %s", id))
	}
}


func getContract(msg proto.Message) Event_Types {
	switch e := msg.(type) {
	case *ProjectCreated:
		return Event_ProjectCreated
	case *DatasetCreated:
		return Event_DatasetCreated
	default:
		panic(errors.Errorf("Uknown event %s", e))

	}
}

func Unmarshal(types Event_Types, buf []byte) proto.Message{
	msg := getInstance(types);
	err := proto.Unmarshal(buf, msg)
	if err != nil {
		panic(errors.Wrap(err, "Failed to unmarshal"))
	}
	return msg

}

func Marshal(msg proto.Message) (Event_Types, []byte){
	id := getContract(msg)
	bytes, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return id, bytes
}
