package logic

import (
	"bookstore/api/internal/svc"
	"bookstore/api/internal/types"
	"bookstore/rpc/add/adder"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (resp *types.AddResp, err error) {
	//// 初始化 TracerProvider
	//provider := otel.GetTracerProvider()
	//
	//tracer := provider.Tracer("bookstore-api")
	//
	//// 创建一个新的 span
	//_, span := tracer.Start(l.ctx, "bookstore-book-add-api")
	//defer span.End()
	//
	//// 设置 span 的属性
	//span.SetAttributes(
	//	attribute.String("book", req.Book),
	//	attribute.Int("price", int(req.Price)),
	//)

	// 手动代码开始
	r, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})

	logx.WithContext(l.ctx).Info("add success")
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: r.Ok,
	}, nil
	// 手动代码结束
}
