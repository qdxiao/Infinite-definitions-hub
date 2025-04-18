package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"infni-backend/apps/api/internal/config"
	"infni-backend/apps/api/internal/db"
)

type ServiceContext struct {
	Config config.Config
	Mysql  sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := db.NewMysql(c.MysqlConfig)
	return &ServiceContext{
		Config: c,
		Mysql:  mysql,
	}
}
