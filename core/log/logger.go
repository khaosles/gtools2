package glog

/*
   @File: logger.go
   @Author: khaosles
   @Time: 2023/4/11 22:16
   @Desc:
*/

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/khaosles/gtools2/core/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const ONE_DAY = time.Hour * 24

var prefix string

var Logger *zap.SugaredLogger

func Init(logCfg *config.Logging) {

	prefix = logCfg.Prefix
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     encodeTime,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	})

	levelConsole := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelChoice(logCfg.LevelConsole)
	})

	levelFile := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= levelChoice(logCfg.LevelFile)
	})

	var cores []zapcore.Core
	if logCfg.LogInFile {
		path := logPath(logCfg.Path)
		hook, _ := rotatelogs.New(
			path+"/%Y-%m-%d.log",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(ONE_DAY*logCfg.MaxHistory),
			rotatelogs.WithRotationTime(ONE_DAY),
			rotatelogs.WithClock(rotatelogs.Local),
		)
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(hook), levelFile))
		if logCfg.LogInConsole {
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelConsole))
		}
	} else {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), levelConsole))
	}
	cores = append(cores)
	core := zapcore.NewTee(cores...)
	log := zap.New(core)
	if logCfg.ShowLine {
		log = log.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
	}
	Logger = log.Sugar()
}

// 自定义调用者的位置编码器
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// 获取相对路径
	relPath := getRelativePath(caller.TrimmedPath())

	// 使用相对路径编码
	enc.AppendString(relPath)
}

// 获取相对路径
func getRelativePath(absPath string) string {
	wk, _ := os.Getwd()
	return strings.TrimPrefix(absPath, filepath.Dir(wk))
}

func logPath(path string) string {
	if path == "" {
		wk, _ := os.Getwd()
		path = wk + "/logs"
	}
	_ = os.MkdirAll(path, os.ModePerm)
	return path
}

func encodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(prefix + t.In(time.FixedZone("CTS", 8*3600)).Format("2006-01-02 15:04:05.000"))
}

func levelChoice(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugf(fmt string, args ...interface{}) {
	Logger.Debugf(fmt, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infof(fmt string, args ...interface{}) {
	Logger.Infof(fmt, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Warnf(fmt string, args ...interface{}) {
	Logger.Warnf(fmt, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorf(fmt string, args ...interface{}) {
	Logger.Errorf(fmt, args...)
}

func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

func Panicf(fmt string, args ...interface{}) {
	Logger.Panicf(fmt, args...)
}
