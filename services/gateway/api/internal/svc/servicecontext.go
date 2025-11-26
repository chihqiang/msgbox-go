// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"chihqiang/msgbox-go/services/common/models"
	"chihqiang/msgbox-go/services/gateway/api/internal/config"
	"chihqiang/msgbox-go/services/gateway/api/internal/middleware"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
	"os"
)

type ServiceContext struct {
	Config              config.Config
	DB                  *gorm.DB
	BasicAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := models.Connect(c.DB)
	if err != nil {
		logx.Errorf("Database connection failed! Error: %v", err)
		os.Exit(1)
	}
	return &ServiceContext{
		Config:              c,
		DB:                  db,
		BasicAuthMiddleware: middleware.NewBasicAuthMiddleware().Handle,
	}
}
