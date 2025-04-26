package svc

import (
	"infni-backend/apps/api/internal/config"
	"infni-backend/apps/api/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	sqlConn := sqlx.NewMysql(c.MysqlConfig.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlConn),
	}
}
