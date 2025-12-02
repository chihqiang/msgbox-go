// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"chihqiang/msgbox-go/pkg/cryptox"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"errors"
	"slices"

	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) error {
	if !slices.Contains([]string{"msgbox-go"}, req.Code) {
		return errors.New("注册码错误")
	}
	var agent models.Agent
	l.svcCtx.DB.Model(&agent).Where(models.Agent{Email: req.Email}).First(&agent)
	if agent.ID > 0 {
		return errors.New("邮箱已注册")
	}
	l.svcCtx.DB.Save(&models.Agent{
		Name:     req.Email,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: cryptox.HashMake(req.Password),
		Status:   true,
	})
	return nil
}
