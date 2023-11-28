package svc

import (
	"article_demo/rpc/internal/config"
	"article_demo/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	UsersModel             model.UsersModel
	ArticlesModel          model.ArticlesModel
	UserSubscriptionsModel model.UserSubscriptionsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                 c,
		UsersModel:             model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
		ArticlesModel:          model.NewArticlesModel(sqlx.NewMysql(c.DataSource), c.Cache),
		UserSubscriptionsModel: model.NewUserSubscriptionsModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
