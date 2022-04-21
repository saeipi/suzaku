package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"suzaku/pkg/constant"
	"time"
	"math/rand"
)
/*
func GetRootPath() string {
	path, _ := filepath.Abs("./")
	reg := regexp.MustCompile(global.ProjectName + "(.*)")
	return reg.ReplaceAllString(path, global.ProjectName+"/")
}
 */

// Determine whether the given path is a folder
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Determine whether the given path is a file
func IsFile(path string) bool {
	return !IsDir(path)
}

func Mkdir(path string) (err error) {
	// 先创建文件夹
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}
	// 再修改权限
	err = os.Chmod(path, os.ModePerm)
	return
}
/*
func FolderMkdir(basePath string, folder string) (folderPath string) {
	environment := single_system.Shared().Env.Name
	folderPath = fmt.Sprintf("%s/%s/", basePath, folder)
	if environment == "" || environment == global.EnvironmentDev {
		folderPath = fmt.Sprintf("./upload/%s/", folder)
	}
	if IsDir(folderPath)==false {
		Mkdir(folderPath)
	}
	return
}
 */

func ReadJson(path string, model interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &model)
	return err
}

func ReadExcel(filePath string) (rows [][]string) {
	rows = ReadXlsx(filePath, "")
	return
}

func ReadXlsx(filePath string, sheet string) (rows [][]string) {
	if sheet == "" {
		sheet = "Sheet1"
	}
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		return
	}
	rows = xlsx.GetRows(sheet)
	return
}

func RemoveFiles(absPath string) (err error) {
	absPath, _ = filepath.Abs(absPath)
	err = filepath.Walk(absPath, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		name := fi.Name()
		p := filepath.Dir(path) + "/" + name
		os.Remove(p)
		return nil
	})
	return
}

// 增量写入文件
func WriteFile(contents string, filePath string) (err error) {
	if contents == "" {
		return
	}
	var f *os.File
	f, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(contents)
	return
}

func GetNewFileNameAndContentType(fileName string, fileType int) (string, string) {
	suffix := path.Ext(fileName)
	newName := fmt.Sprintf("%d-%d%s", time.Now().UnixNano(), rand.Int(), fileName)
	contentType := ""
	if fileType == constant.ImageType {
		contentType = "image/" + suffix[1:]
	}
	return newName, contentType
}
