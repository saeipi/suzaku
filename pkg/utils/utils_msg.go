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

func GenMsgIncr(userID string) string {
	return userID + "_" + int64ToString(GetCurrentTimestampByNano())
}

func GetSessionId(userId1 string, userId2 string) string {
	if userId1 > userId2 {
		return MD5(userId1 + "_" + userId2)
	}
	return MD5(userId2 + "_" + userId1)
}
