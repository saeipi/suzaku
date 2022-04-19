package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
)

func CreateSheet(sheet string, titles []string) (f *excelize.File, columns []string) {
	f = excelize.NewFile()
	f.DeleteSheet(sheet)
	var letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, t := range titles {
		//圈
		var laps = i / 26
		//每圈的下标
		var index = i % 26
		var letter = letters[index]
		var axis = ""
		if laps == 0 {
			axis = letter
		} else {
			axis = letters[laps-1]
			axis += letter
		}
		columns = append(columns, axis)
		f.SetCellValue(sheet, axis+"1", t)
	}
	return
}
