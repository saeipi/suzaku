package config

type Config struct {
	Server Server `yaml:"server"`
	Etcd   Etcd   `yaml:"etcd"`
}
type Server struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}
type Etcd struct {
	Address      []string `yaml:"address"`
	Schema       string   `yaml:"schema"`
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
	DialTimeout  int      `yaml:"dial_timeout"`
}
