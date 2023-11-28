package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"user/user_api/internal/config"
	"user/user_rpc/userrpc"
)

type ServiceContext struct {
	Config config.Config
	User   userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userrpc.NewUserRpc(zrpc.MustNewClient(c.User)),
	}
}
