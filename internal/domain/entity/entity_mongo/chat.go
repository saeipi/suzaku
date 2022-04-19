package entity_mongo

type UserChat struct {
	UID string `bson:"uid"`
	Msg []MessageInfo
}

type MessageInfo struct {
	SendTime int64
	Msg      []byte
}
