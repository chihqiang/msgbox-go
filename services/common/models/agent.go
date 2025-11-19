package models

import (
	"chihqiang/msgbox-go/pkg/cryptox"
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	AgentNo   string         `gorm:"column:agent_no;uniqueIndex;size:32;default:'';comment:编号" json:"agent_no"`
	AgentKey  string         `gorm:"column:agent_key;uniqueIndex;size:32;default:'';comment:agent_key" json:"agent_key"`
	Name      string         `gorm:"column:name;size:32;default:'';comment:联系人姓名" json:"name"`
	Phone     string         `gorm:"column:phone;size:20;default:'';comment:手机号" json:"phone"`
	Email     string         `gorm:"column:email;uniqueIndex;size:100;not null;comment:邮箱" json:"email"`
	Password  string         `gorm:"column:password;size:128;not null;comment:登录密码" json:"-"`
	Status    bool           `gorm:"column:status;not null;default:true;comment:账号状态（true=正常启用，false=禁用/冻结）" json:"status"`
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a Agent) TableName() string {
	return "msgbox_agents"
}

func (a *Agent) VerifyPassword(inputPwd string) bool {
	return cryptox.HashCheck(inputPwd, a.Password)
}
