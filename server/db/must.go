package db

import "github.com/golang/protobuf/proto"

func mustUnmarshal(bytes []byte, msg proto.Message){
	if err := proto.Unmarshal(bytes, msg); err != nil {
		panic("Failed to unmarhsal")
	}

}

