package logger

import "go.uber.org/zap/zapcore"

// LogSection 存放日志配置类
type LogSection struct {
	LogType   string
	LogLevel  string
	FileName  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
	Env       string
}

type LogSetting interface {
	GetLogWriter() zapcore.WriteSyncer
	SetLogLevel() *zapcore.Level
	SetLogEncoder() zapcore.Encoder
}
