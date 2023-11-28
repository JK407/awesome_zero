package logic

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"article_demo/rpc/article"
	"article_demo/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubscribedArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubscribedArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubscribedArticleLogic {
	return &GetSubscribedArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSubscribedArticleLogic) GetSubscribedArticle(in *article.Empty) (*article.GetSubscribedArticleResponse, error) {
	// 获取当前的 span
	span := trace.SpanFromContext(l.ctx)

	// defer 语句确保在函数返回之前结束 span
	defer span.End()
	r, err := l.svcCtx.ArticlesModel.GetSubscribedArticles(l.ctx)
	if err != nil {
		// 记录自定义错误信息并添加到日志中
		l.Logger.Error("GetSubscribedArticleRpcError:", err.Error())
		// 添加错误信息到当前 span 中
		span.SetAttributes(attribute.String("GetSubscribedArticleRpcError:", err.Error()), attribute.Bool("error", true))
	}
	//fmt.Println("GetSubscribedArticleRpc:", r)
	// 将从 GetSubscribedArticles 获取的数据转换成 GetSubscribedArticleResponse 返回
	var subscribedArticles []*article.GetSubscribedArticle
	for _, v := range r {
		subscribedArticles = append(subscribedArticles, &article.GetSubscribedArticle{
			ArticleId:      int32(v.ArticleId),
			ArticleTitle:   v.Title,
			ArticleContent: v.Content,
			ArticleAuthor:  v.Author,
		})
	}
	//fmt.Println("subscribedArticles:", subscribedArticles)
	resp := &article.GetSubscribedArticleResponse{
		Code: 200,
		Msg:  "success",
		Data: subscribedArticles,
	}

	return resp, nil
}
