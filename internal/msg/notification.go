package msg

type NotificationMsg struct {
	SendID      string
	RecvID      string
	Content     []byte //  pb_ws.TipsComm
	MsgFrom     int32
	ContentType int32
	SessionType int32
	OperationID string
}
