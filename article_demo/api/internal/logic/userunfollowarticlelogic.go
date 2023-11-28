package logic

import (
	"article_demo/rpc/articleservice"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"article_demo/api/internal/svc"
	"article_demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUnFollowArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUnFollowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUnFollowArticleLogic {
	return &UserUnFollowArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUnFollowArticleLogic) UserUnFollowArticle(req *types.UserUnFollowArticleRequest) (resp *types.UserUnFollowArticleResponse, err error) {
	// todo: add your logic here and delete this line
	// 获取当前的 span
	span := trace.SpanFromContext(l.ctx)
	// defer 语句确保在函数返回之前结束 span
	defer span.End()
	// 调用 rpc 接口
	r, err := l.svcCtx.ArticleService.UserUnFollowArticle(l.ctx, &articleservice.UserUnFollowArticleRequest{
		UserId:    int32(req.UserId),
		ArticleId: int32(req.ArticleId),
	})
	if err != nil {
		l.Logger.Error("UserUnFollowArticleApiError", attribute.String("err", err.Error()))
		// 添加错误信息到当前 span 中
		span.SetAttributes(attribute.String("UserUnFollowArticleApiError", err.Error()), attribute.Bool("error", true))
	}
	fmt.Println("UserUnFollowArticleApi:", r)
	return
}
