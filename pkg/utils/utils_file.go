package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"suzaku/pkg/constant"
	"time"
)

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

// ----------- i18n -------------
var (
	commonBaseSearchPaths = []string{
		".",
		"..",
		"../..",
		"../../..",
	}
)

func findPath(path string, baseSearchPaths []string, workingDirFirst bool, filter func(os.FileInfo) bool) string {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path
		}

		return ""
	}

	searchPaths := []string{}
	if workingDirFirst {
		searchPaths = append(searchPaths, baseSearchPaths...)
	}

	// Attempt to search relative to the location of the running binary either before
	// or after searching relative to the working directory, depending on `workingDirFirst`.
	var binaryDir string
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.EvalSymlinks(exe); err == nil {
			if exe, err = filepath.Abs(exe); err == nil {
				binaryDir = filepath.Dir(exe)
			}
		}
	}
	if binaryDir != "" {
		for _, baseSearchPath := range baseSearchPaths {
			searchPaths = append(
				searchPaths,
				filepath.Join(binaryDir, baseSearchPath),
			)
		}
	}

	if !workingDirFirst {
		searchPaths = append(searchPaths, baseSearchPaths...)
	}

	for _, parent := range searchPaths {
		found, err := filepath.Abs(filepath.Join(parent, path))
		if err != nil {
			continue
		} else if fileInfo, err := os.Stat(found); err == nil {
			if filter != nil {
				if filter(fileInfo) {
					return found
				}
			} else {
				return found
			}
		}
	}

	return ""
}

func FindPath(path string, baseSearchPaths []string, filter func(os.FileInfo) bool) string {
	return findPath(path, baseSearchPaths, true, filter)
}

// FindFile looks for the given file in nearby ancestors relative to the current working
// directory as well as the directory of the executable.
func FindFile(path string) string {
	return FindPath(path, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir()
	})
}

// fileutils.FindDir looks for the given directory in nearby ancestors relative to the current working
// directory as well as the directory of the executable, falling back to `./` if not found.
func FindDir(dir string) (string, bool) {
	found := FindPath(dir, commonBaseSearchPaths, func(fileInfo os.FileInfo) bool {
		return fileInfo.IsDir()
	})
	if found == "" {
		return "./", false
	}

	return found, true
}

// FindDirRelBinary looks for the given directory in nearby ancestors relative to the
// directory of the executable, then relative to the working directory, falling back to `./` if not found.
func FindDirRelBinary(dir string) (string, bool) {
	found := findPath(dir, commonBaseSearchPaths, false, func(fileInfo os.FileInfo) bool {
		return fileInfo.IsDir()
	})
	if found == "" {
		return "./", false
	}
	return found, true
}
