package events

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"log"
)




// getInstance returns an empty
func getInstance(id Type) proto.Message{
	switch id {
	case Type_Event_ProjectCreated:
		return &ProjectCreated{}
	case Type_Event_DatasetCreated:
		return &DatasetCreated{}
	case Type_Event_DatasetUpdated:
		return &DatasetUpdated{}
	case Type_Event_ExpertAdded:
		return  &ExpertAdded{}
	case Type_Event_JobAdded:
		return &JobAdded{}
	default:
		log.Fatalf("Unknown event type %s", id)
		return nil
	}
}


func GetContract(msg proto.Message) Type {
	switch e := msg.(type) {
	case *ProjectCreated:
		return Type_Event_ProjectCreated
	case *DatasetCreated:
		return Type_Event_DatasetCreated
	case *DatasetUpdated:
		return Type_Event_DatasetUpdated
	case *JobAdded:
		return Type_Event_JobAdded
	case *ExpertAdded:
		return Type_Event_ExpertAdded
	default:
		log.Fatalf("Uknown event %s", e)
		return Type_None

	}
}

func Unmarshal(types Type, buf []byte) proto.Message{
	msg := getInstance(types);
	err := proto.Unmarshal(buf, msg)
	if err != nil {
		panic(errors.Wrap(err, "Failed to unmarshal"))
	}
	return msg

}

func Marshal(msg proto.Message) (Type, []byte){
	id := GetContract(msg)
	bytes, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return id, bytes
}
