package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// 程序配置
type Config struct {
	Etcd    *Etcd    `yaml:"etcd"`
	Mongodb *Mongodb `yaml:"mongodb"`
}
type Etcd struct {
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
	Endpoints    []string `yaml:"endpoints"`
	DialTimeout  int      `yaml:"dial_timeout"`
}
type Mongodb struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
}

var (
	SG_CFG Config
)

func InitConfig(filename string) (err error) {
	var (
		buf []byte
	)
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	yaml.Unmarshal(buf, &SG_CFG)
	return
}
