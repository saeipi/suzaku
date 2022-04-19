package utils

import (
	"errors"
	"strconv"
)

// ToString 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func ToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case string:
		key = value.(string)
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func ToInt(value interface{}) (int, error) {
	switch value.(type) {
	case int:
		return value.(int), nil
	case float64:
		return int(value.(float64)), nil
	case string:
		return strconv.Atoi(value.(string))
	}
	return 0, errors.New("value error")
}

func ToFloat(value interface{}) (float64, error) {
	switch value.(type) {
	case float64:
		return value.(float64), nil
	case int:
		return float64(value.(int)), nil
	case string:
		return strconv.ParseFloat(value.(string), 64)
	}
	return 0, errors.New("value error")
}

func TryToInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch value.(type) {
	case int:
		return value.(int)
	case int32:
		return int(value.(int32))
	case int64:
		return int(value.(int64))
	case float64:
		fValue := value.(float64)
		iValue := int(fValue)
		if float64(iValue) == fValue {
			return iValue
		}
	case float32:
		fValue := value.(float32)
		iValue := int(fValue)
		if float32(iValue) == fValue {
			return iValue
		}
	case string:
		iValue, _ := strconv.Atoi(value.(string))
		return iValue
	}
	return 0
}
