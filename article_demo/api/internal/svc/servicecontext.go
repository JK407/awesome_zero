package svc

import (
	"article_demo/api/internal/config"
	"article_demo/rpc/articleservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	ArticleService articleservice.ArticleService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ArticleService: articleservice.NewArticleService(zrpc.MustNewClient(c.Article)),
	}
}
