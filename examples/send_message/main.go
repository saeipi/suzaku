package main

import (
	"fmt"
	"suzaku/internal/interface/dto/dto_api"
	"suzaku/pkg/utils"
)

func main() {
	msg := dto_api.SendMsgReq{}
	jsStr, _ := utils.ObjToJson(msg)
	fmt.Println(jsStr)
}
