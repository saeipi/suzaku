package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var err error
	var user = UserInfo{
		UserName:    "ayumi",
		UserId:      1003,
		JobNumber:   "KS901",
		Email:       "ksert@163.com",
		PhoneNumber: "17098899839",
		CreatedAt:   time.Now(),
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(user)
	if err != nil {
		return
	}
	map1 := make(map[string]interface{})
	map1["user_name"] = "ayumi"
	//map1["UserId"] = 1003
	//map1["JobNumber"] = "KS901"
	map1["email"] = "ksert@163.com"
	//map1["PhoneNumber"] = "17098899839"
	//map1["CreatedAt"] = "2022-03-18 18:20:30"
	str, err := json.Marshal(map1)

	start := time.Now().UnixNano()
	m := make(map[string]interface{})
	//byt := []byte(`{"user_name":"ayumi","email":"ksert@163.com"}`)
	//var u UserInfo
	for i := 0; i < 100000; i++ {
		//json.Unmarshal(byt, &u)
		//
		//println("")

		//var val UserInfo
		//dec := gob.NewDecoder(&buf)
		//err = dec.Decode(&val)
		json.Unmarshal(str, &m)
	}
	end := time.Now().UnixNano()
	fmt.Println("总运行时长:", (end-start)/1e6, "毫秒")

}

type UserInfo struct {
	UserName    string    `json:"user_name"`    // 姓名
	UserId      int       `json:"user_id"`      // 用户ID
	JobNumber   string    `json:"job_number"`   // 工号
	Email       string    `json:"email"`        // E-Mail
	PhoneNumber string    `json:"phone_number"` // 手机
	CreatedAt   time.Time `json:"created_at"`
}
