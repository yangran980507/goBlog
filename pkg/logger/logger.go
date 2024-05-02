// Package logger 初始化日志
package logger

import (
	"blog/global"
	"blog/pkg/config"
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LogSection struct {
	*config.LogSection
}

type LogSetting interface {
	GetLogWriter() zapcore.WriteSyncer
	SetLogLevel() *zapcore.Level
	SetLogEncoder() zapcore.Encoder
}

// GetLogWriter 获取日志写入介质
func (ls *LogSection) GetLogWriter() zapcore.WriteSyncer {

	// 如果配置了 "daily" ,按日期记录日志文件
	if ls.LogType == "daily" {
		logName := time.Now().Format("2006-01-02")
		ls.FileName = strings.ReplaceAll(ls.FileName,
			"logs.log", logName)
	}

	// 滚动日志配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   ls.FileName,
		MaxSize:    ls.MaxSize,
		MaxAge:     ls.MaxAge,
		MaxBackups: ls.MaxBackup,
	}

	// 返回日志写入介质
	// 本地环境输出日志到文件和终端
	if global.AppSetting.Env == "local" {
		return zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(lumberJackLogger), zapcore.AddSync(os.Stdout))
	} else {
		// 其他环境输出日志到文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// SetLogLevel 设置日志等级
func (ls *LogSection) SetLogLevel() *zapcore.Level {
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(ls.LogLevel)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。")
	}
	return logLevel
}

// SetLogEncoder 设置日志存储格
func (ls *LogSection) SetLogEncoder() zapcore.Encoder {
	encodeConf := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "stacktrace",
		SkipLineEnding: false,
		LineEnding:     zapcore.DefaultLineEnding,
		// 日志级别名称大写,高亮
		EncodeLevel: zapcore.CapitalColorLevelEncoder,

		//时间格式："2006-01-02 15:04:05"
		EncodeTime: customTimeEncoder,

		//以秒为单位
		EncodeDuration: zapcore.SecondsDurationEncoder,

		//短格式路径
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encodeConf)
}

// 设置时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
