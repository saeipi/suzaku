package log

import (
	"bufio"
	"fmt"
	"os"
	"suzaku/pkg/common/config"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *Logger

type Logger struct {
	*logrus.Logger
	Pid int
}

func NewLogger(moduleName string, file string) {
	logger = loggerInit(moduleName, file)
}

func loggerInit(moduleName string, file string) *Logger {
	var (
		src    *os.File
		writer *bufio.Writer
		hook   logrus.Hook
		err    error
	)
	var logger = logrus.New()
	// All logs will be printed
	logger.SetLevel(logrus.Level(config.Config.Log.Level))
	// Close std console output
	src, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	writer = bufio.NewWriter(src)
	logger.SetOutput(writer)
	// logger.SetOutput(os.Stdout)
	// Log Console Print Style Setting
	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})
	//File name and line number display hook
	logger.AddHook(newFileHook())

	//Send logs to elasticsearch hook
	if config.Config.Log.EsSwitch == true {
		logger.AddHook(newEsHook(moduleName))
	}
	//Log file segmentation hook
	hook = NewLfsHook(time.Duration(config.Config.Log.RotationTime)*time.Hour, uint(config.Config.Log.RotationCount), moduleName)
	logger.AddHook(hook)
	return &Logger{
		logger,
		os.Getpid(),
	}
}

func NewLfsHook(rotationTime time.Duration, maxRemainNum uint, moduleName string) logrus.Hook {
	var (
		lfsHook logrus.Hook
	)
	lfsHook = lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.InfoLevel:  initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.WarnLevel:  initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
		logrus.ErrorLevel: initRotateLogs(rotationTime, maxRemainNum, "all", moduleName),
	}, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HideKeys:        false,
		FieldsOrder:     []string{"PID", "FilePath", "OperationID"},
	})
	return lfsHook
}

func initRotateLogs(rotationTime time.Duration, maxRemainNum uint, level string, moduleName string) (writer *rotatelogs.RotateLogs) {
	var (
		err error
	)
	if moduleName != "" {
		moduleName = moduleName + "."
	}
	writer, err = rotatelogs.New(
		config.Config.Log.StorageLocation+moduleName+level+"."+"%Y-%m-%d",
		rotatelogs.WithRotationTime(rotationTime),
		rotatelogs.WithRotationCount(maxRemainNum),
	)
	if err != nil {
		panic(err.Error())
		return
	}
	return
}

func Info(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Infoln(args)
}

func Error(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Errorln(args)
}

func Debug(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Debugln(args)
}

//Deprecated
func Warning(token, OperationID, format string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"PID":         logger.Pid,
		"OperationID": OperationID,
	}).Warningf(format, args...)

}

//Deprecated
func InfoByArgs(format string, args ...interface{}) {
	logger.WithFields(logrus.Fields{}).Infof(format, args)
}

//Deprecated
func ErrorByArgs(format string, args ...interface{}) {
	logger.WithFields(logrus.Fields{}).Errorf(format, args...)
}

//Print log information in k, v format,
//kv is best to appear in pairs. tipInfo is the log prompt information for printing,
//and kv is the key and value for printing.
//Deprecated
func InfoByKv(tipInfo, OperationID string, args ...interface{}) {
	fields := make(logrus.Fields)
	argsHandle(OperationID, fields, args)
	logger.WithFields(fields).Info(tipInfo)
}

//Deprecated
func ErrorByKv(tipInfo, OperationID string, args ...interface{}) {
	fields := make(logrus.Fields)
	argsHandle(OperationID, fields, args)
	logger.WithFields(fields).Error(tipInfo)
}

//Deprecated
func DebugByKv(tipInfo, OperationID string, args ...interface{}) {
	fields := make(logrus.Fields)
	argsHandle(OperationID, fields, args)
	logger.WithFields(fields).Debug(tipInfo)
}

//Deprecated
func WarnByKv(tipInfo, OperationID string, args ...interface{}) {
	fields := make(logrus.Fields)
	argsHandle(OperationID, fields, args)
	logger.WithFields(fields).Warn(tipInfo)
}

//internal method
func argsHandle(OperationID string, fields logrus.Fields, args []interface{}) {
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			fields[fmt.Sprintf("%v", args[i])] = args[i+1]
		} else {
			fields[fmt.Sprintf("%v", args[i])] = ""
		}
	}
	fields["OperationID"] = OperationID
	fields["PID"] = logger.Pid
}

func NewInfo(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Infoln(args)
}

func NewError(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Errorln(args)
}

func NewDebug(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Debugln(args)
}

func NewWarn(OperationID string, args ...interface{}) {
	logger.WithFields(logrus.Fields{
		"OperationID": OperationID,
		"PID":         logger.Pid,
	}).Warnln(args)
}