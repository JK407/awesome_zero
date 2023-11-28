// Code generated by goctl. DO NOT EDIT.
// Source: article.proto

package server

import (
	"context"

	"article_demo/rpc/article"
	"article_demo/rpc/internal/logic"
	"article_demo/rpc/internal/svc"
)

type ArticleServiceServer struct {
	svcCtx *svc.ServiceContext
	article.UnimplementedArticleServiceServer
}

func NewArticleServiceServer(svcCtx *svc.ServiceContext) *ArticleServiceServer {
	return &ArticleServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServiceServer) GetSubscribedArticle(ctx context.Context, in *article.Empty) (*article.GetSubscribedArticleResponse, error) {
	l := logic.NewGetSubscribedArticleLogic(ctx, s.svcCtx)
	return l.GetSubscribedArticle(in)
}

func (s *ArticleServiceServer) UserFollowArticle(ctx context.Context, in *article.UserFollowArticleRequest) (*article.UserFollowArticleResponse, error) {
	l := logic.NewUserFollowArticleLogic(ctx, s.svcCtx)
	return l.UserFollowArticle(in)
}

func (s *ArticleServiceServer) UserUnFollowArticle(ctx context.Context, in *article.UserUnFollowArticleRequest) (*article.UserUnFollowArticleResponse, error) {
	l := logic.NewUserUnFollowArticleLogic(ctx, s.svcCtx)
	return l.UserUnFollowArticle(in)
}
