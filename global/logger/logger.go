// Package logger 存放 logger 对象
package logger

import (
	"blog/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitializeLog() {
	setupLogger()
}

func setupLogger() {
	//获取日志写入介质
	writeSyncer := global.LogSetting.GetLogWriter()

	//设置日志级别
	level := global.LogSetting.SetLogLevel()

	//设置日志存储格式
	encoder := global.LogSetting.SetLogEncoder()

	//初始化 core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	//初始化 Logger
	Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	//替换全局 Logger
	zap.ReplaceGlobals(Logger)
}
