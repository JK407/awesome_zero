package logic

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"strconv"
	"user/user_rpc/internal/svc"
	"user/user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.Request) (*user.Response, error) {
	// 获取当前的 span
	span := trace.SpanFromContext(l.ctx)

	// defer 语句确保在函数返回之前结束 span
	defer span.End()

	// 查询用户信息
	r, err := l.svcCtx.Model.FindOne(l.ctx, in.Id)
	if err != nil {
		// 记录自定义错误信息并添加到日志中
		l.Logger.Error("查询失败", attribute.String("err", err.Error()))
		// 添加错误信息到当前 span 中
		span.SetAttributes(attribute.String("get_user_by_id", "根据Id查询用户失败"), attribute.Bool("error", true), attribute.String("err", err.Error()))
		return nil, err
	}
	l.Logger.Infof("查询成功", attribute.String("id", strconv.FormatInt(in.Id, 10)))
	// 构建响应
	response := &user.Response{
		Code: 200,
		Msg:  "success",
		Data: &user.User{
			Id:          r.Id,
			Name:        r.Name,
			Description: r.Description,
		},
	}
	return response, nil
}
