package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"unsafe"
)

func Int2Bytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func Bytes2Int(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BufferDecode(buf []byte, obj interface{}) (err error) {
	var (
		buffer  *bytes.Buffer
		decoder *gob.Decoder
	)
	buffer = bytes.NewBuffer(buf)
	decoder = gob.NewDecoder(buffer)
	err = decoder.Decode(obj)
	return
}

func ObjEncode(obj interface{}) (buf []byte, err error) {
	var (
		buffer  bytes.Buffer
		encoder *gob.Encoder
	)
	encoder = gob.NewEncoder(&buffer)
	err = encoder.Encode(obj)
	if err != nil {
		return
	}
	buf = buffer.Bytes()
	return
}
