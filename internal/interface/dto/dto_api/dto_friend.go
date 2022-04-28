package dto_api

type AddFriendReq struct {
	FromUserId string `json:"from_user_id"`
	ToUserId   string `json:"to_user_id" binding:"required"`
}
