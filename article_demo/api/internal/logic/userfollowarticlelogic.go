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

type UserFollowArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowArticleLogic {
	return &UserFollowArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowArticleLogic) UserFollowArticle(req *types.UserFollowArticleRequest) (resp *types.UserFollowArticleResponse, err error) {
	// todo: add your logic here and delete this line
	// 获取当前的 span
	span := trace.SpanFromContext(l.ctx)

	// defer 语句确保在函数返回之前结束 span
	defer span.End()
	// 调用 rpc 接口
	r, err := l.svcCtx.ArticleService.UserFollowArticle(l.ctx, &articleservice.UserFollowArticleRequest{
		UserId:    int32(req.UserId),
		ArticleId: int32(req.ArticleId),
	})
	if err != nil {
		l.Logger.Error("UserFollowArticleApiError", attribute.String("err", err.Error()))
		// 添加错误信息到当前 span 中
		span.SetAttributes(attribute.String("UserFollowArticleApiError", err.Error()), attribute.Bool("error", true))
		return nil, err
	}
	fmt.Println("UserFollowArticleApi:", r)
	resp = &types.UserFollowArticleResponse{
		Code: int(r.Code),
		Msg:  r.Msg,
		Data: types.UserFollowArticle{
			SubscriptionId: int(r.Data.SubscriptionId),
		},
	}
	return resp, nil
}
