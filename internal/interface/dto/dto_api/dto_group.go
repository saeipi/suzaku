package dto_api

type GroupMemberListReq struct {
	PageReq
	GroupId     string `json:"group_id"`
	OperationId string `json:"operation_id"`
}

type GroupMemberListResp struct {
	TotalRows  int64              `json:"total_rows"`
	MemberList []*GroupMemberInfo `json:"member_list"`
}

type GroupMemberInfo struct {
	GroupId       string `json:"group_id"`
	UserId        string `json:"user_id"`
	Nickname      string `json:"nickname"`
	UserAvatarUrl string `json:"user_avatar_url"`
	RoleLevel     int32  `json:"role_level"`
	//JoinTime       int64  `json:"join_time"`
	//JoinSource     int32  `json:"join_source"`
	//OperatorUserId string `json:"operator_user_id"`
	MuteEndTime int64 `json:"mute_end_time"`
	//Ex          string `json:"ex"`
}
