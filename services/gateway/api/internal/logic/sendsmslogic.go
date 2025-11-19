// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"chihqiang/msgbox-go/services/common/errors"
	"chihqiang/msgbox-go/services/gateway/api/internal/svc"
	"chihqiang/msgbox-go/services/gateway/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
	"time"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendRequest) (resp *types.SmsResponse, err error) {
	username, ok := l.ctx.Value(types.BasicAuthUsername).(string)
	if !ok || username == "" {
		l.Logger.Errorf("SendSms missing valid username from ctx")
		return nil, errors.ErrAuthInvalid
	}
	password, ok := l.ctx.Value(types.BasicAuthPassword).(string)
	if !ok || password == "" {
		l.Logger.Errorf("SendSms missing valid password from ctx, username: %s", username)
		return nil, errors.ErrAuthInvalid
	}
	return &types.SmsResponse{
		TraceID:      trace.TraceIDFromContext(l.ctx),
		MsgID:        1,
		FailCount:    2,
		SuccessCount: 0,
		Time:         time.Now().UnixMicro(),
	}, nil
}
