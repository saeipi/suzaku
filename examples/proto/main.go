package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	pb_auth "suzaku/pkg/proto/auth"
	"unsafe"
)

func main() {
	var buf01 = int2bytes(10001)
	var buf02 = str2bytes("49ba59abbe56e057")

	var m1 = &pb_auth.CommonResp{
		Code: 1,
		Msg:  "测试",
	}

	buf03, err := proto.Marshal(m1)
	if err != nil {
		fmt.Println(err.Error())
	}

	var m2 pb_auth.CommonResp
	proto.Unmarshal(buf03, &m2)

	var buffer bytes.Buffer // Buffer是一个实现了读写方法的可变大小的字节缓冲
	buffer.Write(buf01)     // 4个长度的消息类型
	buffer.Write(buf02)     // 16个长度的签名
	buffer.Write(buf03)     // 数据

	buf04 := buffer.Bytes()
	buf05 := buf04[20:len(buf04)]

	var m3 pb_auth.CommonResp
	if err = proto.Unmarshal(buf05, &m3); err != nil {
		fmt.Println(err.Error())
	}

	var file = "/Volumes/data/proto/byte.txt"
	if err = ioutil.WriteFile(file, buf05, 0666); err != nil {
		fmt.Println(err.Error())
	}
	var m4 pb_auth.CommonResp
	var buf06 []byte
	buf06, err = ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = proto.Unmarshal(buf06, &m4); err != nil {
		fmt.Println(err.Error())
	}

	var in int
	fmt.Scan(&in)
}

func int2bytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func bytes2int(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
