package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/service/demo/api/internal/middleware"
	"go-zero-demo/service/demo/dao"

	"go-zero-demo/service/demo/api/internal/config"
	"go-zero-demo/service/demo/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *dao.Query
	AccessLog rest.Middleware
	DbEngin   *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	//conn := sqlx.NewMysql(c.Mysql.DataSource)
	//启动Gorm支持
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.PGSql.DataSource,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:    c,
		DbEngin:   db,
		UserModel: dao.Use(db),
		//UserModel: model.NewUserModel(conn, c.RedisCache),
		AccessLog: middleware.NewAccessLogMiddleware().Handle,
	}
}
