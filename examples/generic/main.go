package main

import (
	"fmt"
	"suzaku/pkg/utils"
)

// 像声明接口一样声明
type MySlice interface {
	[]int | []int8 | []int16 | []int32 | []string
}

// T的类型为声明的MySlice
func GetSlicesLen[T MySlice](a, b T, num int) int {
	return len(a) + len(b) + num
}

func main() {
	v := GetSlicesLen([]string{"1"}, []string{"1"}, 5)
	fmt.Println(v)
	ExportExcel(make([]RowInfo1, 0), make([]ColumnField, 0), "", make([]string, 0))
}

type ColumnField struct {
	Key   string `json:"key" form:"key"`
	Type  string `json:"type" form:"type"` //string,int,float,percent,time
	Title string `json:"title" form:"title"`
}

type RowInfo1 struct {
	Key   string `json:"key" form:"key"`
	Type  string `json:"type" form:"type"` //string,int,float,percent,time
	Title string `json:"title" form:"title"`
}

type RowInfo2 struct {
	Key   string `json:"key" form:"key"`
	Type  string `json:"type" form:"type"` //string,int,float,percent,time
	Title string `json:"title" form:"title"`
}

type ExcelColumnList interface {
	RowInfo1 | RowInfo2
}

func ExportExcel[T ExcelColumnList](list []T, columns []ColumnField, sheet string, titles []string) {
	for _, v := range list {
		if utils.ToString(v) != "" {

		}
	}
}
