package utils

import (
	"math/rand"
	"strconv"
)

func GetMsgID(sendID string) string {
	t := int64ToString(GetCurrentTimestampByNano())
	return MD5(t + sendID + int64ToString(rand.Int63n(GetCurrentTimestampByNano())))
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
