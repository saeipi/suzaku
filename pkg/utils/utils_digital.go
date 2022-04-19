package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math"
	"strconv"
)

// 取两个切片的交集
func IntIntersect(slice1 []int, slice2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

// 取要校验的和已经校验过的差集，找出需要校验的切片IP（找出slice1中 slice2中没有的）
func IntDifference(slice1, slice2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	inter := IntIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			n = append(n, v)
		}
	}
	return n
}

func FloatThousandSeperator(val float64, point int) (res string) {
	printer := message.NewPrinter(language.English)
	key := "%." + strconv.Itoa(point) + "f"
	res = printer.Sprintf(key, val)
	return
}

func SymbolThousandSeperator(val float64, point int, symbol string) (res string) {
	var isNegative bool
	if val < 0 {
		isNegative = true
		val = math.Abs(val)
	}
	printer := message.NewPrinter(language.English)
	key := "%s%." + strconv.Itoa(point) + "f"
	res = printer.Sprintf(key, symbol, val)
	if isNegative {
		res = "-" + res
	}
	return
}

func IntThousandSeperator(val int) (res string) {
	printer := message.NewPrinter(language.English)
	res = printer.Sprintf("%d", val)
	return
}
