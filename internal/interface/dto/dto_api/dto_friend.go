package dto_api

type AddFriendReq struct {
	OperationId string `json:"operation_id" binding:"max=40"`
	FromUserId  string `json:"from_user_id" binding:"max=40"`
	ToSzkId     string `json:"to_szk_id" binding:"required,max=40"`
	ReqMsg      string `json:"req_msg" binding:"required,max=255"` // 添加好友消息
}

type FriendRequestListReq struct {
	UserId   string `json:"user_id"`
	Role     int    `form:"role" json:"role" binding:"required,oneof=0 1 2"`
	Page     int    `form:"page" json:"page" binding:"required,min=1"`
	PageSize int    `form:"page_size" json:"page_size" binding:"required,min=10,max=20"`
}

type FriendRequestListResp struct {
	TotalRows int                  `json:"total_rows"`
	List      []*FriendRequestItem `json:"list"`
}

type FriendRequestItem struct {
	ReqId        string `json:"req_id"`        // 事件ID
	FromUserId   string `json:"from_user_id"`  // 发起人ID
	ToUserId     string `json:"to_user_id"`    // 目标人ID
	ReqMsg       string `json:"req_msg"`       // 添加好友消息
	HandleResult int32  `json:"handle_result"` // 结果
	HandleMsg    string `json:"handle_msg"`    // 处理消息
	HandledTs    int64  `json:"handled_ts"`
	ReqTs        int64  `json:"req_ts"` // 请求时间
}

type HandleFriendRequestReq struct {
	FromUserId   string `json:"from_user_id" binding:"required,max=40"`       // 发起人ID
	UserId       string `json:"user_id" binding:"max=40"`                     // 目标人ID 处理人ID
	HandleUserId string `json:"handle_user_id" binding:"max=40"`              // 处理人ID
	HandleMsg    string `json:"handle_msg" binding:"max=255"`                 // 处理消息
	HandleResult int32  `json:"handle_result" binding:"required,min=1,max=2"` // 结果
}
