package logic

import (
	"context"
	"user/user_rpc/userrpc"

	"user/user_api/internal/svc"
	"user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.GetUserByIdRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// 手动代码开始
	r, _ := l.svcCtx.User.GetUserById(l.ctx, &userrpc.Request{
		Id: req.Id,
	})
	//fmt.Println(err)
	return &types.Response{
		Code: int(r.Code),
		Msg:  r.Msg,
		Data: r.Data,
	}, nil
	// 手动代码结束
	return
}
