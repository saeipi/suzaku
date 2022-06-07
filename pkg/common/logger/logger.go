package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"suzaku/pkg/common/config"
	"time"
)

const (
	LowercaseLevelEncoder      = "Lowercase"
	LowercaseColorLevelEncoder = "LowercaseColor"
	CapitalLevelEncoder        = "Capital"
	CapitalColorLevelEncoder   = "CapitalColor"
)

var (
	logger *zap.SugaredLogger
)

func InitLogger(cfg *config.Zap) {
	// 判断是否有Director文件夹
	directory := "./logs/" + cfg.Directory
	if _, err := os.Stat(directory); err != nil {
		_ = os.Mkdir(directory, os.ModePerm)
	}

	// zap.LevelEnablerFunc(func(lev zapcore.Level) bool 用来划分不同级别的输出
	// 根据不同的级别输出到不同的日志文件

	// 调试级别
	debugLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.DebugLevel
	})
	// 日志级别
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.InfoLevel
	})
	// 警告级别
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel
	})
	// 错误级别
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.ErrorLevel
	})
	// panic级别
	panicLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.DPanicLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("%s/debug.log", directory), debugLevel, cfg),
		getEncoderCore(fmt.Sprintf("%s/info.log", directory), infoLevel, cfg),
		getEncoderCore(fmt.Sprintf("%s/warn.log", directory), warnLevel, cfg),
		getEncoderCore(fmt.Sprintf("%s/error.log", directory), errorLevel, cfg),
		getEncoderCore(fmt.Sprintf("%s/panic.log", directory), panicLevel, cfg),
	}

	//zapcore.NewTee(cores ...zapcore.Core) zapcore.Core
	//NewTee创建一个Core，将日志条目复制到两个或更多的底层Core中

	log := zap.New(zapcore.NewTee(cores[:]...))
	//用文件名、行号和zap调用者的函数名注释每条消息
	if cfg.ShowLine {
		log = log.WithOptions(zap.AddCaller())
	}
	logger = log.Sugar()
	logger.Sync()
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(filename string, level zapcore.LevelEnabler, cfg *config.Zap) (core zapcore.Core) {
	// 使用lumberjack进行日志分割
	writer := getWriteSyncer(filename, cfg)
	return zapcore.NewCore(getEncoder(cfg), writer, level)
}

func getWriteSyncer(filename string, cfg *config.Zap) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   filename,               // 日志文件的位置
		MaxSize:    cfg.Segment.MaxSize,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: cfg.Segment.MaxBackups, // 保留旧文件的最大个数
		MaxAge:     cfg.Segment.MaxBackups, // 保留旧文件的最大天数
		Compress:   cfg.Segment.Compress,   // 是否压缩/归档旧文件
	}
	if cfg.LogStdout == true {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	}
	return zapcore.AddSync(hook)
}

// getEncoder 获取zapcore.Encoder
func getEncoder(cfg *config.Zap) zapcore.Encoder {
	// 获取配置文件的输出格式 json or console
	switch cfg.Encoder {
	case "json":
		return zapcore.NewJSONEncoder(getEncoderConfig(cfg))
	case "console":
		return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(cfg *config.Zap) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: cfg.StacktraceKey,         // 栈名
		LineEnding:    zapcore.DefaultLineEnding, // 默认的结尾\n
		//EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写字母输出 zapcore.LowercaseLevelEncoder
		EncodeTime:     customTimeEncoder,              // 时间格式 zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.SecondsDurationEncoder, // 编码间隔
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 绝对路径:zapcore.FullCallerEncoder,相对路径:zapcore.ShortCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 根据配置文件重新配置编码颜色和字体
	switch cfg.EncodeLevel {
	case LowercaseLevelEncoder: // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case LowercaseColorLevelEncoder: // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case CapitalLevelEncoder: // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case CapitalColorLevelEncoder: // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
}

// 自定义日志级别显示
func customEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// 自定义行号显示
func customEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.Info(err)
	}
	return dir
}

func GetFilePath() string {
	logfile := GetCurrentDirectory() + "/" + GetAppname() + ".log"
	return logfile
}

func GetAppname() string {
	full := os.Args[0]
	splits := strings.Split(full, "/")
	if len(splits) >= 1 {
		name := splits[len(splits)-1]
		return name
	}
	return ""
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger.Errorw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger.Panicw(msg, keysAndValues...)
}
