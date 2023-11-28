package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSubscriptionsModel = (*customUserSubscriptionsModel)(nil)

type (
	// UserSubscriptionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSubscriptionsModel.
	UserSubscriptionsModel interface {
		userSubscriptionsModel
	}

	customUserSubscriptionsModel struct {
		*defaultUserSubscriptionsModel
	}
)

// NewUserSubscriptionsModel returns a model for the database table.
func NewUserSubscriptionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSubscriptionsModel {
	return &customUserSubscriptionsModel{
		defaultUserSubscriptionsModel: newUserSubscriptionsModel(conn, c, opts...),
	}
}
