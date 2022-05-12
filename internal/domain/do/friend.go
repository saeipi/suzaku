package do

type FriendInfo struct {
	SessionId string `json:"session_id"` // 会话ID
	UserId    string `json:"user_id"`    // 用户ID
	SzkId     string `json:"szk_id"`     // 账户ID
	Nickname  string `json:"nickname"`   // 昵称
	Gender    string `json:"gender"`     // 性别
	AvatarUrl string `json:"avatar_url"` // 头像
}
