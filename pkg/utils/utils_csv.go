package utils

import (
	"encoding/csv"
	"errors"
	"os"
)

// 导出Csv
func ExportCsv(headers []string, rows [][]string, filePath string) error {
	// 1、创建文件
	newFile, err := os.Create(filePath)
	if err != nil {
		return errors.New("创建文件失败")
	}
	defer newFile.Close()

	// 2、写入UTF-8
	// newFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	// 3、写数据到csv文件
	w := csv.NewWriter(newFile)
	// 4、WriteAll方法使用Write方法向w写入多条记录，并在最后调用Flush方法清空缓存。
	datas := make([][]string, 0)
	if len(headers) > 0 {
		datas = append(datas, headers)
	}
	datas = append(datas, rows...)
	w.WriteAll(datas)
	w.Flush()
	return err
}
