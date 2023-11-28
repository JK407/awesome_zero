package logic

import (
	"context"

	"article_demo/rpc/article"
	"article_demo/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUnFollowArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUnFollowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUnFollowArticleLogic {
	return &UserUnFollowArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUnFollowArticleLogic) UserUnFollowArticle(in *article.UserUnFollowArticleRequest) (*article.UserUnFollowArticleResponse, error) {
	// todo: add your logic here and delete this line

	return &article.UserUnFollowArticleResponse{}, nil
}
