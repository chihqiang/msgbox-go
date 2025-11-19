// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"chihqiang/msgbox-go/services/gateway/api/internal/config"
	"chihqiang/msgbox-go/services/gateway/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	BasicAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		BasicAuthMiddleware: middleware.NewBasicAuthMiddleware().Handle,
	}
}
