// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package auth

import (
	"chihqiang/msgbox-go/pkg/cryptox"
	"chihqiang/msgbox-go/services/agent/api/internal/svc"
	"chihqiang/msgbox-go/services/agent/api/internal/types"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var agent models.Agent
	err = l.svcCtx.DB.Model(models.Agent{}).Where("email = ?", req.Email).First(&agent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logger.Errorf("login failed: phone=%s not found", req.Email)
			return nil, errors.New("账号不存在")
		}
		return nil, err
	}
	// 4. 校验账号状态（禁用状态无法登录）
	if !agent.Status {
		l.Logger.Errorf("login failed: email=%s is disabled", req.Email)
		return nil, errors.New("账号已禁用，请联系平台管理员")
	}
	// 5. 密码校验（调用模型层 VerifyPassword 方法）
	if !agent.VerifyPassword(req.Password) {
		l.Logger.Errorf("login failed: email=%s,error=%v", req.Email, err)
		return nil, errors.New("登录密码验证错误")
	}
	accessToken, err := l.GenerateAccessToken(agent.ID, agent.Phone)
	if err != nil {
		l.Logger.Errorf("generate access token email=%s,failed: %v", req.Email, err)
		return nil, errors.New("令牌生成失败，请稍后重试")
	}
	return &types.LoginResp{
		ID:        agent.ID,
		Name:      agent.Name,
		Token:     accessToken,
		ExpiresIn: l.svcCtx.Config.Auth.AccessExpire,
	}, nil
}
func (l *LoginLogic) GenerateAccessToken(agentID int64, phone string) (accessToken string, err error) {
	expireTime := time.Now().Add(time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second)
	claims := jwt.MapClaims{
		types.JWTAgentID: agentID,
		types.JWTPhone:   phone,
		"exp":            jwt.NewNumericDate(expireTime),
		"iat":            jwt.NewNumericDate(time.Now()),
		"iss":            "msgbox-api",
	}
	return cryptox.JWTEncode(l.svcCtx.Config.Auth.AccessSecret, claims)
}
