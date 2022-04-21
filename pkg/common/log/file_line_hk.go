package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type fileHook struct{}

func newFileHook() *fileHook {
	return &fileHook{}
}

func (f *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f *fileHook) Fire(entry *logrus.Entry) error {
	entry.Data["FilePath"] = findCaller(6)
	return nil
}

//func (f *fileHook) Fire(entry *logrus.Entry) error {
//	var s string
//	_, b, c, _ := runtime.Caller(10)
//	i := strings.SplitAfter(b, "/")
//	if len(i) > 3 {
//		s = i[len(i)-3] + i[len(i)-2] + i[len(i)-1] + ":" + utils.IntToString(c)
//	}
//	entry.Data["FilePath"] = s
//	return nil
//}

func findCaller(skip int) string {
	var (
		file string
		line int
		i    int
	)
	for i = 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, "log") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	var (
		file string
		line int
		ok   bool
		n    int
		i    int
	)
	_, file, line, ok = runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	for i = len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
