package dto_api

type AddFriendReq struct {
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id" binding:"required"`
	ReqMsg     string `json:"req_msg" binding:"max=32"` // 添加好友消息
}
