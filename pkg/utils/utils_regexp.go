package utils

import (
	"regexp"
	"strconv"
)

const (
	// 获取数值
	RegexpExpressionNumericalValue = "(\\-)?\\d+(\\.\\d+)?$"
	// 验证邮箱
	RegexpExpressionVerifyEmail = "^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\\.){1,4}[a-z]{2,4}$"
	// 验证手机
	RegexpExpressionVerifyMobile = "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	// 验证英文字母
	RegexpExpressionVerifyEnglishAlphabet = "^[A-Za-z]+$"
	// 数字
	RegexpExpressionDigital = "\\d+\\.?\\d*"
	// 验证年月
	RegexpExpressionVerifyYearMonth = "^\\d{4}年\\d{1,2}月"
	// int 数值
	RegexpExpressionIntDigitals = "\\d+"
)

// 验证邮箱
func VerifyEmail(email string) bool {
	reg := regexp.MustCompile(RegexpExpressionVerifyEmail)
	return reg.MatchString(email)
}

// 验证手机
func VerifyMobile(mobileNum string) bool {
	reg := regexp.MustCompile(RegexpExpressionVerifyMobile)
	return reg.MatchString(mobileNum)
}

// 验证英文字母
func VerifyEnglishAlphabet(str string) bool {
	reg := regexp.MustCompile(RegexpExpressionVerifyEnglishAlphabet)
	return reg.MatchString(str)
}

func GetFirstDigital(text string) (digital int, err error) {
	reg := regexp.MustCompile(RegexpExpressionDigital)
	res := reg.FindStringSubmatch(text)
	if len(res) > 0 {
		digital, err = strconv.Atoi(res[0])
		return
	}
	return
}

// 验证年月 ^\d{4}-\d{1,2}-\d{1,2}*/
func VerifyYearMonth(date string) bool {
	if date == "" {
		return false
	}
	//regular := "^\\d{4}(\\u5E74)\\d{1,2}(\\u6708)"
	regular := RegexpExpressionVerifyYearMonth
	reg := regexp.MustCompile(regular)
	return reg.MatchString(date)
}

// 获取年月
func GetYearMonth(date string) (res string) {
	if date == "" {
		return
	}
	reg := regexp.MustCompile(RegexpExpressionVerifyYearMonth)
	strs := reg.FindStringSubmatch(date)
	if len(strs) > 0 {
		res = strs[0]
	}
	return
}

func GetIntDigitals(text string) (digitals []int, err error) {
	reg := regexp.MustCompile(RegexpExpressionIntDigitals)
	textArr := reg.FindAllStringSubmatch(text, -1)
	for _, texts := range textArr {
		for _, t := range texts {
			var digital int
			digital, err = strconv.Atoi(t)
			if err != nil {
				continue
			}
			digitals = append(digitals, digital)
		}
	}
	return
}

func VerifyRegular(text string, regular string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(text)
}

func GetRegularIntDigitals(text string, regular string) (digitals []int, err error) {
	reg := regexp.MustCompile(regular)
	res := reg.FindStringSubmatch(text)
	for _, n := range res {
		var digital int
		digital, err = strconv.Atoi(n)
		if err != nil {
			return
		}
		digitals = append(digitals, digital)
	}
	return
}

// 获取UTC偏移值
func UtcOffset(text string) (offset float64) {
	var (
		err  error
		reg  *regexp.Regexp
		list []string
	)
	reg = regexp.MustCompile(RegexpExpressionNumericalValue)
	list = reg.FindStringSubmatch(text)
	for _, str := range list {
		offset, err = strconv.ParseFloat(str, 64)
		if err != nil {
			continue
		}
		offset *= 3600
		return
	}
	if offset == 0 {
		offset = 8 * 3600
	}
	return
}
