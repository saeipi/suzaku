package utils

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func SqlIntIn(values []int) (res string) {
	var inValue = "("
	for i, v := range values {
		if i == len(values)-1 {
			inValue = inValue + strconv.Itoa(v) + ")"
		} else {
			inValue = inValue + strconv.Itoa(v) + ","
		}
	}
	return inValue
}

func OrderField(field string, values []string, sort string) (res string) {
	if field == "" || len(values) == 0 {
		return
	}
	res = fmt.Sprintf("FIELD(`%s`", field)
	for _, v := range values {
		res = res + "," + fmt.Sprintf("'%s'", v)
	}
	res = res + ") " + sort
	return
}

// GetBranchInsertSql 获取批量添加数据sql语句
func GetBranchInsertSql(objs []interface{}, tableName string) string {
	if len(objs) == 0 {
		return ""
	}
	fieldName := ""
	var valueTypeList []string
	fieldNum := reflect.TypeOf(objs[0]).NumField()
	fieldType := reflect.TypeOf(objs[0])
	for i := 0; i < fieldNum; i++ {
		name := GetColumnName(fieldType.Field(i).Tag.Get("gorm"))
		// 添加字段名
		if i == fieldNum-1 {
			fieldName += fmt.Sprintf("`%s`", name)
		} else {
			fieldName += fmt.Sprintf("`%s`,", name)
		}
		// 获取字段类型
		fieldType := fieldType.Field(i).Type.Name()
		valueTypeList = append(valueTypeList, fieldType)
	}
	var valueList []string
	for _, obj := range objs {
		objVal := reflect.ValueOf(obj)
		val := "("
		for index, i := range valueTypeList {
			if index == fieldNum-1 {
				val += GetFormatField(objVal, index, i, "")
			} else {
				val += GetFormatField(objVal, index, i, ",")
			}
		}
		val += ")"
		valueList = append(valueList, val)
	}
	insertSql := fmt.Sprintf("insert into `%s` (%s) values %s", tableName, fieldName, strings.Join(valueList, ",")+";")
	return insertSql
}

// GetFormatField 获取字段类型值转为字符串
func GetFormatField(objVal reflect.Value, index int, fieldType string, sep string) (val string) {
	switch fieldType {
	case "string":
		val += fmt.Sprintf("'%s'%s", objVal.Field(index).String(), sep)
	case "Time":
		var t = objVal.Field(index).Interface().(time.Time)
		val += fmt.Sprintf("'%s'%s", t.Format(ExtsTimeStandard), sep)
	case "int":
		val += fmt.Sprintf("%d%s", objVal.Field(index).Int(), sep)
	case "uint":
		val += fmt.Sprintf("%d%s", objVal.Field(index).Uint(), sep)
	case "bool":
		val += fmt.Sprintf("%d%s", objVal.Field(index).Int(), sep)
	case "float64":
		val += fmt.Sprintf("%f%s", objVal.Field(index).Float(), sep)
	}
	return val
}

// GetColumnName 获取字段名
func GetColumnName(jsonName string) string {
	for _, name := range strings.Split(jsonName, ";") {
		if strings.Index(name, "column") == -1 {
			continue
		}
		return strings.Replace(name, "column:", "", 1)
	}
	return ""
}

// BatchCreateModelsByPage 分页批量插入
func BatchCreateModelsByPage(tx *gorm.DB, dataList []interface{}, tableName string) (err error) {
	if len(dataList) == 0 {
		return
	}
	// 如果超过一百条, 则分批插入
	size := 100
	page := len(dataList) / size
	if len(dataList)%size != 0 {
		page += 1
	}
	for i := 1; i <= page; i++ {
		var bills = make([]interface{}, 0)
		if i == page {
			bills = dataList[(i-1)*size:]
		} else {
			bills = dataList[(i-1)*size : i*size]
		}
		sql := GetBranchInsertSql(bills, tableName)
		if err = tx.Exec(sql).Error; err != nil {
			fmt.Println(fmt.Sprintf("batch create data error: %v, sql: %s, tableName: %s", err, sql, tableName))
			return
		}
	}
	return
}
