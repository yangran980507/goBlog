package config

import "time"

// ServerSection 存放服务配置类
type ServerSection struct {
	HttpPort     string
	RunMode      string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

// MysqlSection 存放 mysql 配置类
type MysqlSection struct {
	Host            string
	Port            string
	DBName          string
	UserName        string
	Password        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifeTime time.Duration
}

// AppSection 存放 App 配置类
type AppSection struct {
	Name             string
	Env              string
	URL              string
	JWTSecretKey     string
	JWTExpireTime    int64
	JWTMaxExpireTime int64
}

// LogSection 存放日志配置类
type LogSection struct {
	LogType   string
	LogLevel  string
	FileName  string
	MaxSize   int
	MaxBackup int
	MaxAge    int
}

// RedisSection 存放 redis 配置
type RedisSection struct {
	//
}
