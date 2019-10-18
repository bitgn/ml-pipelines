package db

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func mustUnmarshal(bytes []byte, msg proto.Message){
	if err := proto.Unmarshal(bytes, msg); err != nil {
		panic("Failed to unmarhsal")
	}

}



func mustUid(uid []byte){
	if len(uid)==0{
		log.Panicln("UID can't be nil")
	}
}

func mustName(name string){
	if len(name)==0{
		log.Panicln("Name can't be nil")
	}
}

