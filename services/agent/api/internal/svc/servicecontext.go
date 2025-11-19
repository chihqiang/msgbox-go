// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"chihqiang/msgbox-go/services/agent/api/internal/config"
	"chihqiang/msgbox-go/services/common/models"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"os"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := models.Connect(c.DB)
	if err != nil {
		logc.Errorf(context.Background(), "Database connection failed! Error: %v", err)
		os.Exit(1)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
