package utils

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func Proto2Map(pb proto.Message) (pbMap map[string]interface{}, err error) {
	var (
		buffer     bytes.Buffer
		marshaller *jsonpb.Marshaler
		buf        []byte
	)
	buffer = bytes.Buffer{}
	marshaller = &jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  true,
		EmitDefaults: true,
	}
	_ = marshaller.Marshal(&buffer, pb)
	buf = buffer.Bytes()
	err = json.Unmarshal(buf, &pbMap)
	return
}
