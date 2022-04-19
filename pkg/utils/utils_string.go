package utils

import (
	"bytes"
	"encoding/hex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strconv"
	"strings"
	"suzaku/pkg/constant"
)

func ParseFloat(str string) (res float64) {
	if str == "" {
		return
	}
	res, _ = strconv.ParseFloat(str, 64)
	return
}

func ParseInt(str string) (res int) {
	if str == "" {
		return
	}
	res, _ = strconv.Atoi(str)
	return
}

func SqlStringIn(values []string) (res string) {
	res = "(" + strings.Join(values, ",") + ")"
	return
}

func StrToIntArray(str string, sep string) (res []int) {
	arr := strings.Split(str, sep)
	for _, v := range arr {
		res = append(res, ParseInt(v))
	}
	return
}

// 获取中文字符串第一个首字母
func ChinesesLetter(chinese string) (res string) {
	for _, c := range chinese {
		res = res + ChineseFirstLetter(string(c))
	}
	return
}

// 获取中文字符串第一个首字母
func ChineseFirstLetter(chinese string) string {
	// 获取中文字符串第一个字符
	firstChar := string([]rune(chinese)[:1])

	// Utf8 转 GBK2312
	firstCharGbk, err := Utf8ToGbk([]byte(firstChar))
	if err != nil {
		return ""
	}

	// 获取第一个字符的16进制
	firstCharHex := hex.EncodeToString(firstCharGbk)

	// 16进制转十进制
	firstCharDec, err := strconv.ParseInt(firstCharHex, 16, 0)
	if err != nil {
		return ""
	}

	// 十进制落在GB 2312的某个拼音区间即为某个字母
	firstCharDecimalRelative := firstCharDec - 65536
	if firstCharDecimalRelative >= -20319 && firstCharDecimalRelative <= -20284 {
		return "A"
	}
	if firstCharDecimalRelative >= -20283 && firstCharDecimalRelative <= -19776 {
		return "B"
	}
	if firstCharDecimalRelative >= -19775 && firstCharDecimalRelative <= -19219 {
		return "C"
	}
	if firstCharDecimalRelative >= -19218 && firstCharDecimalRelative <= -18711 {
		return "D"
	}
	if firstCharDecimalRelative >= -18710 && firstCharDecimalRelative <= -18527 {
		return "E"
	}
	if firstCharDecimalRelative >= -18526 && firstCharDecimalRelative <= -18240 {
		return "F"
	}
	if firstCharDecimalRelative >= -18239 && firstCharDecimalRelative <= -17923 {
		return "G"
	}
	if firstCharDecimalRelative >= -17922 && firstCharDecimalRelative <= -17418 {
		return "H"
	}
	if firstCharDecimalRelative >= -17417 && firstCharDecimalRelative <= -16475 {
		return "J"
	}
	if firstCharDecimalRelative >= -16474 && firstCharDecimalRelative <= -16213 {
		return "K"
	}
	if firstCharDecimalRelative >= -16212 && firstCharDecimalRelative <= -15641 {
		return "L"
	}
	if firstCharDecimalRelative >= -15640 && firstCharDecimalRelative <= -15166 {
		return "M"
	}
	if firstCharDecimalRelative >= -15165 && firstCharDecimalRelative <= -14923 {
		return "N"
	}
	if firstCharDecimalRelative >= -14922 && firstCharDecimalRelative <= -14915 {
		return "O"
	}
	if firstCharDecimalRelative >= -14914 && firstCharDecimalRelative <= -14631 {
		return "P"
	}
	if firstCharDecimalRelative >= -14630 && firstCharDecimalRelative <= -14150 {
		return "Q"
	}
	if firstCharDecimalRelative >= -14149 && firstCharDecimalRelative <= -14091 {
		return "R"
	}
	if firstCharDecimalRelative >= -14090 && firstCharDecimalRelative <= -13319 {
		return "S"
	}
	if firstCharDecimalRelative >= -13318 && firstCharDecimalRelative <= -12839 {
		return "T"
	}
	if firstCharDecimalRelative >= -12838 && firstCharDecimalRelative <= -12557 {
		return "W"
	}
	if firstCharDecimalRelative >= -12556 && firstCharDecimalRelative <= -11848 {
		return "X"
	}
	if firstCharDecimalRelative >= -11847 && firstCharDecimalRelative <= -11056 {
		return "Y"
	}
	if firstCharDecimalRelative >= -11055 && firstCharDecimalRelative <= -10247 {
		return "Z"
	}
	return ""
}

// Utf8ToGbk
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GetConversationIDBySessionType(sourceID string, sessionType int) string {
	switch sessionType {
	case constant.SingleChatType:
		return "single_" + sourceID
	case constant.GroupChatType:
		return "group_" + sourceID
	}
	return ""
}
