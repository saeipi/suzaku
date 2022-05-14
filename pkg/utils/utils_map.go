package utils

import jsoniter "github.com/json-iterator/go"

func MergeStringMaps(maps ...map[string]interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return
}

func MergeIntMaps(maps ...map[int]int) (result map[int]int) {
	result = make(map[int]int)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return
}

func GetSwitchFromOptions(Options map[string]bool, key string) (result bool) {
	if flag, ok := Options[key]; !ok || flag {
		return true
	}
	return false
}

func SetSwitchFromOptions(options map[string]bool, key string, value bool) {
	if options == nil {
		options = make(map[string]bool, 5)
	}
	options[key] = value
}

func IntMapToInterfaceMap(intMap map[string]int) (result map[string]interface{}) {
	result = map[string]interface{}{}
	for k, v := range intMap {
		result[k] = v
	}
	return
}

func StringMapToIntMap(strMap map[string]string) (result map[string]int) {
	result = map[string]int{}
	for k, v := range strMap {
		result[k] = TryToInt(v)
	}
	return
}

func MapToStruct(in map[string]string, out interface{}) (err error) {
	//err = mapstructure.Decode(in, out)
	var (
		buf []byte
	)
	buf, err = jsoniter.Marshal(in)
	if err != nil {
		return
	}
	err = jsoniter.Unmarshal(buf, out)
	return
}
