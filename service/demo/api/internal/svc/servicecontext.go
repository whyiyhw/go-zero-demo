package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"

	"go-zero-demo/service/demo/api/internal/config"
	"go-zero-demo/service/demo/api/internal/middleware"
	"go-zero-demo/service/demo/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	AccessLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.RedisCache),
		AccessLog: middleware.NewAccessLogMiddleware().Handle,
	}
}
