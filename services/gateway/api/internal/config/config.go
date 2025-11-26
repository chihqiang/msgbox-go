// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"chihqiang/msgbox-go/services/common/models"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB models.Config
}
