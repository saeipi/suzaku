package dto_api

type AddFriendReq struct {
	UserID   string `json:"_"`
	TargetUserId string `json:"target_user_id" binding:"required"`
}
