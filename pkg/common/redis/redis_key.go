package redis

const (
	RedisKeyJwtUserToken = "token:%s:%d"        // 用户TokenKey:userId-platform
	RedisKeyGroup        = "group:%s"           // 群ID
	RedisKeyGroupMember  = "group_member:%s:%s" // 群ID:用户ID
)
