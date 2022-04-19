package utils

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

func YamlToStruct(file string, target *interface{}) (err error) {
	var content []byte
	content, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &target)
	return
}
