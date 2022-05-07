package rpc_chat

import (
	"math/rand"
	"strconv"
	"suzaku/pkg/utils"
	"time"
)

/*
	now:= time.Now()
	fmt.Println(now.Unix()) // 1565084298 秒
	fmt.Println(now.UnixNano()) // 1565084298178502600 纳秒
	fmt.Println(now.UnixNano() / 1e6) // 1565084298178 毫秒
*/
func GetMsgID(sendID string) string {
	return utils.MD5(strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + sendID + "-" + strconv.Itoa(rand.Int()))
}
