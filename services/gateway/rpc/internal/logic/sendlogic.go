package logic

import (
	"context"

	"chihqiang/msgbox-go/services/gateway/rpc/internal/svc"
	"chihqiang/msgbox-go/services/gateway/rpc/pb/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendLogic) Send(in *rpc.SendRequest) (*rpc.SendResponse, error) {
	// todo: add your logic here and delete this line

	return &rpc.SendResponse{}, nil
}
