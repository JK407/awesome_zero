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

type GetSubscribedArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSubscribedArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribedArticleLogic {
	return &GetSubscribedArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSubscribedArticleLogic) GetSubscribedArticle() (resp *types.GetSubscribedArticleResponse, err error) {
	// todo: add your logic here and delete this line
	// 获取当前的 span
	span := trace.SpanFromContext(l.ctx)

	// defer 语句确保在函数返回之前结束 span
	defer span.End()
	r, err := l.svcCtx.ArticleService.GetSubscribedArticle(l.ctx, &articleservice.Empty{})
	if err != nil {
		// 记录自定义错误信息并添加到日志中
		l.Logger.Error("GetSubscribedArticleApiError", attribute.String("err", err.Error()))
		// 添加错误信息到当前 span 中
		span.SetAttributes(attribute.String("GetSubscribedArticleApiError", err.Error()), attribute.Bool("error", true))
	}
	fmt.Println("GetSubscribedArticleApi:", r)
	var subscribedArticles []types.GetSubscribedArticle
	for _, v := range r.Data {
		subscribedArticles = append(subscribedArticles, types.GetSubscribedArticle{
			ArticleId:      int(v.ArticleId),
			ArticleTitle:   v.ArticleTitle,
			ArticleContent: v.ArticleContent,
			ArticleAuthor:  v.ArticleAuthor,
		})
	}
	resp = &types.GetSubscribedArticleResponse{
		Code: int(r.Code),
		Msg:  r.Msg,
		Data: subscribedArticles,
	}
	return
}
