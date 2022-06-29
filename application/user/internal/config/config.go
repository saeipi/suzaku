package config

type Config struct {
	Name      string    `yaml:"name"`
	RPCServer RPCServer `yaml:"rpc_server"`
}

type RPCServer struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}
