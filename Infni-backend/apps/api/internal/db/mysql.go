package db

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"infni-backend/apps/api/internal/config"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)
	db, err := mysql.RawDB()
	if err != nil {
		panic(fmt.Errorf("数据库连接失败，原因：:%v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(mysqlConfig.ConnectTimeout))
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		panic(fmt.Errorf("数据库 Ping 失败，原因：:%v", err))
	}

	db.SetMaxOpenConns(mysqlConfig.MaxOpenConns)
	db.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	return mysql
}
