package utils

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ObjToJson(obj interface{}) (str string, err error) {
	var (
		buf []byte
	)
	buf, err = json.Marshal(obj)
	if err != nil {
		return
	}
	str = string(buf)
	return
}

func TryObjToJson(obj interface{}) (str string) {
	str, _ = ObjToJson(obj)
	return
}

func JsonToObj(str string, obj interface{}) error {
	return json.Unmarshal([]byte(str), obj)
}

func TryJsonToObj(str string, obj interface{}) {
	_ = JsonToObj(str, obj)
}

func JsonToMap(str string) (maps map[string]interface{}) {
	json.Unmarshal([]byte(str), &maps)
	return
}

func MergeJsonList(list []string) (jsStr string) {
	if len(list) == 0 {
		return
	}
	builder := strings.Builder{}
	jsStr = strings.Join(list, ",")
	builder.WriteString("[")
	builder.WriteString(jsStr)
	builder.WriteString("]")
	jsStr = builder.String()
	return
}
