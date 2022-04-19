package redis

import (
	"suzaku/pkg/constant"
	"suzaku/pkg/utils"
)

const (
	AccountTempCode               = "ACCOUNT_TEMP_CODE"
	resetPwdTempCode              = "RESET_PWD_TEMP_CODE"
	userIncrSeq                   = "REDIS_USER_INCR_SEQ:" // user incr seq
	appleDeviceToken              = "DEVICE_TOKEN"
	userMinSeq                    = "REDIS_USER_MIN_SEQ:"
	uidPidToken                   = "UID_PID_TOKEN_STATUS:"
	conversationReceiveMessageOpt = "CON_RECV_MSG_OPT:"
)

func JudgeAccountEXISTS(account string) (bool, error) {
	key := AccountTempCode + account
	return KeyExists(key)
}
func SetAccountCode(account string, code, ttl int) (err error) {
	key := AccountTempCode + account
	Set(key, code, ttl)
	return err
}
func GetAccountCode(account string) (string, error) {
	key := AccountTempCode + account
	return Get(key)
}

//Perform seq auto-increment operation of user messages
func IncrUserSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	return GetUint64(key)
}

//Get the largest Seq
func GetUserMaxSeq(uid string) (uint64, error) {
	key := userIncrSeq + uid
	return GetUint64(key)
}

//Set the user's minimum seq
func SetUserMinSeq(uid string, minSeq uint32) (err error) {
	key := userMinSeq + uid
	err = Set(key, minSeq, 0)
	return err
}

//Get the smallest Seq
func GetUserMinSeq(uid string) (uint64, error) {
	key := userMinSeq + uid
	return GetUint64(key)
}

//Store Apple's device token to redis
func SetAppleDeviceToken(accountAddress, value string) (err error) {
	key := appleDeviceToken + accountAddress
	err = Set(key, value, 0)
	return err
}

//Delete Apple device token
func DelAppleDeviceToken(accountAddress string) (err error) {
	key := appleDeviceToken + accountAddress
	err = Del(key)
	return
}

//Store userid and platform class to redis
func AddTokenFlag(userID string, platformID int32, token string, flag int) (err error) {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	err = HSet(key, token, flag)
	return
}

func GetTokenMapByUidPid(userID, platformID string) (map[string]int, error) {
	key := uidPidToken + userID + ":" + platformID
	return GetIntMap(key)
}

func SetTokenMapByUidPid(userID string, platformID int32, m map[string]int) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	return HMSet(key, utils.IntMapToInterfaceMap(m))
}

func DeleteTokenByUidPid(userID string, platformID int32, fields []string) error {
	key := uidPidToken + userID + ":" + constant.PlatformIDToName(platformID)
	return HDels(key, fields)
}

func SetSingleConversationMsgOpt(userID, conversationID string, opt int) error {
	key := conversationReceiveMessageOpt + userID
	return HSet(key, conversationID, opt)
}

func GetSingleConversationMsgOpt(userID, conversationID string) (int, error) {
	key := conversationReceiveMessageOpt + userID
	return HGetInt(key, conversationID)
}

func GetAllConversationMsgOpt(userID string) (result map[string]int, err error) {
	key := conversationReceiveMessageOpt + userID
	maps := map[string]string{}
	maps, err = HGetAll(key)
	if err != nil {
		result = map[string]int{}
		return
	}
	result = utils.StringMapToIntMap(maps)
	return
}

func SetMultiConversationMsgOpt(userID string, m map[string]int) error {
	key := conversationReceiveMessageOpt + userID
	return HMSet(key, utils.IntMapToInterfaceMap(m))
}

func GetMultiConversationMsgOpt(userID string, conversationIDs []string) (m map[string]int, err error) {
	m = make(map[string]int)
	key := conversationReceiveMessageOpt + userID
	var list []interface{}
	list, err = HMGet(key, conversationIDs...)
	if err != nil {
		return
	}
	for k, v := range conversationIDs {
		m[v] = utils.TryToInt(list[k])
	}
	return
}
