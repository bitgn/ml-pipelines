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
	case Type_Event_DatasetVersionAdded:
		return &DatasetVersionAdded{}
	case Type_Event_JobRunStarted:
		return &JobRunStarted{}
	case Type_Event_JobRunLogged:
		return &JobRunLogged{}
	case Type_Event_JobRunCompleted:
		return &JobRunCompleted{}
	case Type_Event_JobRunFailed:
		return &JobRunFailed{}
	default:
		log.Panicf("Unknown event type %s", id)
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
	case *DatasetVersionAdded:
		return Type_Event_DatasetVersionAdded
	case *JobRunStarted:
		return Type_Event_JobRunStarted
	case *JobRunLogged:
		return Type_Event_JobRunLogged
	case *JobRunFailed:
		return Type_Event_JobRunFailed
	case *JobRunCompleted:
		return Type_Event_JobRunCompleted
	default:
		log.Panicf("Uknown event %T", e)
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


