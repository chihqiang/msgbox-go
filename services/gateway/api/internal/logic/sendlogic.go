// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"chihqiang/msgbox-go/services/common/errs"
	"chihqiang/msgbox-go/services/common/models"
	"chihqiang/msgbox-go/services/common/pipeline"
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"chihqiang/msgbox-go/services/gateway/api/internal/svc"
	"chihqiang/msgbox-go/services/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req *types.SendRequest) (resp *types.SendResponse, err error) {
	username, ok := l.ctx.Value(types.BasicAuthUsername).(string)
	if !ok || username == "" {
		l.Logger.Errorf("Send missing valid username from ctx")
		return nil, errs.ErrAuthInvalid
	}
	password, ok := l.ctx.Value(types.BasicAuthPassword).(string)
	if !ok || password == "" {
		l.Logger.Errorf("Send missing valid password from ctx, username: %s", username)
		return nil, errs.ErrAuthInvalid
	}
	send, err := l.sendPipeline(username, password, req)
	if err != nil {
		l.Logger.Errorf("Send failed, err: %v", err)
		return nil, err
	}
	return &types.SendResponse{
		TraceID:      trace.TraceIDFromContext(l.ctx),
		BatchNo:      send.BatchNo,
		FailCount:    send.FailCount,
		SuccessCount: send.SuccessCount,
		Time:         send.CreatedAt.UnixMicro(),
	}, nil
}

func (l *SendLogic) sendPipeline(agentNo, agentKey string, req *types.SendRequest) (*models.SendBatch, error) {
	sendPipeline := pipeline.SendPipeline{
		DB:           l.svcCtx.DB,
		Log:          l.Logger,
		AgentNo:      agentNo,
		AgentKey:     agentKey,
		TemplateCode: req.TemplateCode,
		Receivers:    req.Receivers,
		Variables:    req.Variables,
		Extra:        req.Extra,
	}
	if err := sendPipeline.Check(l.ctx); err != nil {
		return nil, err
	}
	if err := sendPipeline.Send(l.ctx); err != nil {
		return nil, err
	}
	return sendPipeline.GetSendBatch()
}
