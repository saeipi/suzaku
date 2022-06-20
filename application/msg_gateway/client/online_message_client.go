package client

type OnlineMessageClient interface {
}

type onlineMessageClient struct {
}

func NewOnlineMessageClient() OnlineMessageClient {
	return &onlineMessageClient{}
}
