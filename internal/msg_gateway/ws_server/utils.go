package ws_server

import (
	"bytes"
	"encoding/binary"
)

func bytes2Int(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func int2Bytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func decode(msgCode int, buf []byte) interface{} {
	return nil
}

func encode(msgCode int, body []byte) (buf []byte) {
	var (
		buyCode []byte
		buffer  bytes.Buffer
	)
	buyCode = int2Bytes(msgCode)
	buffer.Write(buyCode)
	if body != nil {
		buffer.Write(body)
	}
	buf = buffer.Bytes()
	return
}
