package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"suzaku/examples/send_msg/client"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/utils"
	"time"
)

func main() {
	recvID := "1524257293623889920"
	recvToken := "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQ4MzgyMzksImlzcyI6InRLaDNmTXdXZXBDZHM5amJUdUo5SUdZSlliUElabmoiLCJvcmlnX2lhdCI6MTY1MjI0NjIzOSwicGxhdGZvcm1faWQiOjEsInVzZXJfaWQiOiIxNTI0MjU3MjkzNjIzODg5OTIwIn0.xF1_6TY88pGh-zC8qc6-AJ0KpGChfQ-kayzdGxW9SXY"
	userId := "1524255191468085248"
	ts := utils.GetCurrentTimestampByMill()

	client.NewClient(recvID, recvToken)

	time.Sleep(time.Second * 2)

	msg := dto_api.SendMsgReq{
		SenderPlatformID: 1,
		SendID:           userId, // 发送者ID
		SenderNickName:   "无敌",
		SenderAvatarUrl:  "https://github.com/saeipi/suzaku/blob/main/assets/images/suzaku.jpg",
		OperationID:      userId,
		Data: dto_api.SendMsgData{
			SessionType: 1,   // 单聊为1，群聊为2
			MsgFrom:     100, // 100:用户消息 200:系统消息
			ContentType: 101, // 消息类型，101表示文本，102表示图片
			SessionId:   utils.GetSessionId(userId, recvID),
			RecvID:      recvID, // 接收者ID
			GroupID:     "",
			ForceList:   nil,
			Content:     nil,
			Options:     nil,
			ClientMsgID: utils.GetMsgID(userId),
			CreatedTs:   ts,
			OffLineInfo: nil,
		},
	}
	fmt.Println("|--------------| 发送消息时间:", time.Now(), "|--------------|")
	msg.Data.Content = utils.Str2Bytes("文本聊天消息 1523642393075650560")
	buf, err := Post("http://127.0.0.1:10000/api/chat/send_msg", msg, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))

	var input int
	fmt.Scan(&input)
}

func Post(url string, data interface{}, timeOutSecond int) (respBuf []byte, err error) {
	var (
		jsonBuf []byte
		req     *http.Request
		client  *http.Client
		resp    *http.Response
	)
	jsonBuf, err = json.Marshal(data)
	if err != nil {
		return
	}
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonBuf))
	if err != nil {
		return
	}
	req.Close = true
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Cookie", "jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQ2OTE2MjksImlzcyI6InRLaDNmTXdXZXBDZHM5amJUdUo5SUdZSlliUElabmoiLCJvcmlnX2lhdCI6MTY1MjA5OTYyOSwicGxhdGZvcm1faWQiOjEsInVzZXJfaWQiOiIxNTIzNjQyMzkzMDc1NjUwNTYwIn0.w6AAE-3S7I4kpIS9FBRyDx7hFZTcBHJcPK_2Isk_3v4")

	client = &http.Client{Timeout: time.Duration(timeOutSecond) * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBuf, err = ioutil.ReadAll(resp.Body)
	return
}
