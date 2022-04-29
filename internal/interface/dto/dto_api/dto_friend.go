package dto_api

type AddFriendReq struct {
	OperationId string `json:"operation_id" binding:"required"`
	FromUserId  string `json:"from_user_id"`
	ToSzkId     string `json:"to_szk_id" binding:"required"`
	ReqMsg      string `json:"req_msg" binding:"max=32"` // 添加好友消息
}
