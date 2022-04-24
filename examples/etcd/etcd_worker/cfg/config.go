package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"suzaku/pkg/common/config"
)

// 程序配置
type Config struct {
	Etcd    EtcdConfig         `yaml:"etcd"`
	Mongodb config.MongoConfig `yaml:"mongodb"`
}

type EtcdConfig struct {
	Port         int      `yaml:"port"`
	ReadTimeout  int      `yaml:"read_timeout"`
	WriteTimeout int      `yaml:"write_timeout"`
	Endpoints    []string `yaml:"endpoints"`
	DialTimeout  int      `yaml:"dial_timeout"`
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
