package config

type values map[string]interface{}

// ServerSection 存放服务配置
//
//	type ServerSection struct {
//		Values map[string]interface{}
//	}
type ServerSection struct {
	HttpPort string
	Env      string
	URL      string
}

// MysqlSection 存放 mysql 配置
type MysqlSection struct {
	Values values
}

// RedisSection 存放 redis 配置
type RedisSection struct {
	Values values
}
