package config

type Config struct {
	Name      string    `yaml:"name"`
	WsServer  WsServer  `yaml:"ws_server"`
	RPCServer RPCServer `yaml:"rpc_server"`
}
type WsServer struct {
	Name                  string `yaml:"name"`
	Port                  int    `yaml:"port"`
	WriteWait             int    `yaml:"write_wait"`
	PongWait              int    `yaml:"pong_wait"`
	PingPeriod            int    `yaml:"ping_period"`
	MaxMessageSize        int    `yaml:"max_message_size"`
	ReadBufferSize        int    `yaml:"read_buffer_size"`
	WriteBufferSize       int    `yaml:"write_buffer_size"`
	HeaderLength          int    `yaml:"header_length"`
	ChanClientSendMessage int    `yaml:"chan_client_send_message"`
	ChanServerReadMessage int    `yaml:"chan_server_read_message"`
	ChanServerRegister    int    `yaml:"chan_server_register"`
	ChanServerUnregister  int    `yaml:"chan_server_unregister"`
	MaxConnections        int    `yaml:"max_connections"`
	MinimumTimeInterval   int    `yaml:"minimum_time_interval"`
}
type RPCServer struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
