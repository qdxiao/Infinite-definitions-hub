package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	Jwt         Jwt
}

type MysqlConfig struct {
	DataSource     string
	MaxOpenConns   int
	MaxIdleConns   int
	ConnectTimeout int
}

type Jwt struct {
	Secret string
	Expire int
}
