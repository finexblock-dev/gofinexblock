package utils

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func MessagesToJson(s proto.Message) ([]byte, error) {
	var marshaller protojson.MarshalOptions
	var jsonData []byte
	var err error

	marshaller = protojson.MarshalOptions{UseProtoNames: true}

	jsonData, err = marshaller.Marshal(s)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}