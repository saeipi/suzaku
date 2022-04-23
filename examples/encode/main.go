package main

import (
	"suzaku/pkg/utils"
	"time"
)

type UserInfo struct {
	UserName    string    `json:"user_name"`    // 姓名
	UserId      int       `json:"user_id"`      // 用户ID
	JobNumber   string    `json:"job_number"`   // 工号
	Email       string    `json:"email"`        // E-Mail
	PhoneNumber string    `json:"phone_number"` // 手机
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	var (
		user    UserInfo
		buf     []byte
		newUser UserInfo
		err     error
	)
	user = UserInfo{
		UserName:    "ayumi",
		UserId:      1003,
		JobNumber:   "KS901",
		Email:       "ksert@163.com",
		PhoneNumber: "17098899839",
		CreatedAt:   time.Now(),
	}
	buf, err = utils.ObjEncode(user)
	if err != nil {
		return
	}
	newUser = UserInfo{}
	err = utils.BufferDecode(buf, &newUser)
	if err != nil {
		return
	}
}
