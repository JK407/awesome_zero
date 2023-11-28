package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/user_rpc/internal/config"
	"user/user_rpc/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
