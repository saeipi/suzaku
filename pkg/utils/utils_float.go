package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

func FormatFloat(num float64, decimal int) (res string) {
	if num == float64(int64(num)) {
		res = fmt.Sprintf("%d", int64(num))
	} else {
		d := math.Pow10(decimal)
		res = strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	}
	return
}

func Decimal2(value float64) (res float64) {
	res, _ = decimal.NewFromFloat(value).Round(2).Float64()
	return
}

func Decimal4(value float64) (res float64) {
	res, _ = decimal.NewFromFloat(value).Round(4).Float64()
	return
}
