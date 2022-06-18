package config

type Config struct {
	Name      string    `yaml:"name"`
	WsServer  WsServer  `yaml:"ws_server"`
	RPCServer RPCServer `yaml:"rpc_server"`
}
type WsServer struct {
	Name            string `yaml:"name"`
	Port            int    `yaml:"port"`
	WriteWait       int    `yaml:"write_wait"`
	PongWait        int    `yaml:"pong_wait"`
	MaxMessageSize  int    `yaml:"max_message_size"`
	ReadBufferSize  int    `yaml:"read_buffer_size"`
	WriteBufferSize int    `yaml:"write_buffer_size"`
}
type RPCServer struct {
	Name         string `yaml:"name"`
	Port         int    `yaml:"port"`
}