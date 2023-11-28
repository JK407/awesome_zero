package logic

import (
	"article_demo/rpc/model"
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"article_demo/rpc/article"
	"article_demo/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFollowArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowArticleLogic {
	return &UserFollowArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserFollowArticleLogic) UserFollowArticle(in *article.UserFollowArticleRequest) (*article.UserFollowArticleResponse, error) {
	span := trace.SpanFromContext(l.ctx)
	defer span.End()
	resp := &article.UserFollowArticleResponse{}
	//  先查询是订阅表中是否有这条订阅记录
	_, findSubscriptionErr := l.svcCtx.UserSubscriptionsModel.FindOneByUserIdAndArticleId(l.ctx, int64(in.UserId), int64(in.ArticleId))
	fmt.Println("findSubscriptionErr:", findSubscriptionErr)
	//findSubscriptionErr: unsupported unmarshal type

	if findSubscriptionErr == nil {
		//  如果有，就返回已经关注过
		l.Logger.Error("UserFollowArticleError:", "已经关注过")
		span.SetAttributes(attribute.String("UserFollowArticleError:", "已经关注过"), attribute.Bool("error", true))
		resp = &article.UserFollowArticleResponse{
			Code: 400,
			Msg:  "已经关注过",
			Data: &article.UserFollowArticle{
				SubscriptionId: 0,
			},
		}
	} else {
		//  如果没有
		//  查询用户是否存在
		_, findUserErr := l.svcCtx.UsersModel.FindOne(l.ctx, int64(in.UserId))
		if findUserErr != nil {
			l.Logger.Error("FindUserByIdError:", findUserErr.Error())
			// 添加错误信息到当前 span 中
			span.SetAttributes(attribute.String("FindUserByIdError:", findUserErr.Error()), attribute.Bool("error", true))
		}
		//  再查询文章是否存在
		if _, findArticleErr := l.svcCtx.ArticlesModel.FindOne(l.ctx, int64(in.ArticleId)); findArticleErr != nil {
			l.Logger.Error("FindArticleByIdError:", findArticleErr.Error())
			// 添加错误信息到当前 span 中
			span.SetAttributes(attribute.String("FindArticleByIdError:", findArticleErr.Error()), attribute.Bool("error", true))
		}
		insertSubscriptionData := &model.UserSubscriptions{
			UserId:    int64(in.UserId),
			ArticleId: int64(in.ArticleId),
		}
		//  如果都存在，就添加关注
		if _, insertSubscriptionErr := l.svcCtx.UserSubscriptionsModel.Insert(l.ctx, insertSubscriptionData); insertSubscriptionErr != nil {
			l.Logger.Error("InsertSubscriptionError:", insertSubscriptionErr.Error())
			span.SetAttributes(attribute.String("InsertSubscriptionError:", insertSubscriptionErr.Error()), attribute.Bool("error", true))
		}
		resp = &article.UserFollowArticleResponse{
			Code: 200,
			Msg:  "success",
			Data: &article.UserFollowArticle{
				SubscriptionId: int32(insertSubscriptionData.SubscriptionId),
			},
		}
	}

	return resp, nil
}
