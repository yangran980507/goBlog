package config

import "time"

// ServerSection 存放服务配置类
type ServerSection struct {
	HttpPort string
	Env      string
	URL      string
}

// MysqlSection 存放 mysql 配置
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

// RedisSection 存放 redis 配置
type RedisSection struct {
	//
}
