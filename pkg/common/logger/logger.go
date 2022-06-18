package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	LowercaseLevelEncoder      = "Lowercase"
	LowercaseColorLevelEncoder = "LowercaseColor"
	CapitalLevelEncoder        = "Capital"
	CapitalColorLevelEncoder   = "CapitalColor"
)

type Zap struct {
	Encoder       string  `json:"encoder" yaml:"encoder"`               // 编码器 console Or json
	Directory     string  `json:"directory"  yaml:"directory"`          // 日志文件夹
	ShowLine      bool    `json:"show_line" yaml:"show_line"`           // 显示行
	EncodeLevel   string  `json:"encode_level" yaml:"encode_level"`     // 编码级
	StacktraceKey string  `json:"stacktrace_key" yaml:"stacktrace_key"` // 栈名
	LogStdout     bool    `json:"log_stdout" yaml:"log_stdout"`         // 输出控制台
	Segment       Segment `json:"segment" yaml:"segment"`               // 日志分割
}

type Segment struct {
	MaxSize    int  `json:"maxsize" yaml:"maxsize"`
	MaxAge     int  `json:"maxage" yaml:"maxage"`
	MaxBackups int  `json:"maxbackups" yaml:"maxbackups"`
	LocalTime  bool `json:"localtime" yaml:"localtime"`
	Compress   bool `json:"compress" yaml:"compress"`
}

type Logger struct {
	logger *zap.SugaredLogger
	cfg    *Zap
}

func NewLogger(cfg *Zap) *Logger {
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

	directory := "./logs/" + cfg.Directory
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
	if cfg.ShowLine == true {
		log = log.WithOptions(zap.AddCaller())
	}
	logger := log.Sugar()
	logger.Sync()

	return &Logger{logger, cfg}
}

func getEncoderCore(filename string, level zapcore.LevelEnabler, cfg *Zap) (core zapcore.Core) {
	// 使用lumberjack进行日志分割
	writer := getWriteSyncer(filename, cfg)
	return zapcore.NewCore(getEncoder(cfg), writer, level)
}

func getWriteSyncer(filename string, cfg *Zap) zapcore.WriteSyncer {
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

func getEncoder(cfg *Zap) zapcore.Encoder {
	switch cfg.Encoder {
	case "json":
		return zapcore.NewJSONEncoder(getEncoderConfig(cfg))
	case "console":
		return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
}

func getEncoderConfig(cfg *Zap) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,              // 栈名
		LineEnding:     zapcore.DefaultLineEnding,      // 默认的结尾\n
		EncodeTime:     customTimeEncoder,              // 时间格式 zapcore.ISO8601TimeEncoder
		EncodeDuration: zapcore.SecondsDurationEncoder, // 编码间隔
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 绝对路径:zapcore.FullCallerEncoder,相对路径:zapcore.ShortCallerEncoder
		EncodeName:     zapcore.FullNameEncoder,
	}
	switch cfg.EncodeLevel {
	case LowercaseLevelEncoder:
		// 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case LowercaseColorLevelEncoder:
		// 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case CapitalLevelEncoder:
		// 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case CapitalColorLevelEncoder:
		// 大写编码器带颜色
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

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}
