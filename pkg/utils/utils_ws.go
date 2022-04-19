package utils

import "strconv"

func WsIdentifier(userID string, platformID int32) string {
	return userID + "-" + strconv.Itoa(int(platformID))
}
