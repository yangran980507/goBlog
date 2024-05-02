// Package logger 存放 logger 对象
package logger

import (
	"blog/global"
	loggerpkg "blog/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitializeLog() {
	setupLogger()
}

func setupLogger() {

	logSetting := loggerpkg.LogSection{
		LogSection: global.LogSetting,
	}
	//获取日志写入介质
	writeSyncer := logSetting.GetLogWriter()

	//设置日志级别
	level := logSetting.SetLogLevel()

	//设置日志存储格式
	encoder := logSetting.SetLogEncoder()

	//初始化 core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	//初始化 Logger
	loggerpkg.Logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	//替换全局 Logger
	zap.ReplaceGlobals(loggerpkg.Logger)
}
