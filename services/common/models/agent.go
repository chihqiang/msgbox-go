package models

import (
	"chihqiang/msgbox-go/pkg/cryptox"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Agent struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	AgentNo     string         `gorm:"column:agent_no;uniqueIndex;size:32;default:'';comment:编号" json:"agent_no"`
	AgentSecret string         `gorm:"column:agent_secret;size:32;default:'';comment:密钥" json:"agent_secret"`
	Name        string         `gorm:"column:name;size:32;default:'';comment:联系人姓名" json:"name"`
	Phone       string         `gorm:"column:phone;size:20;default:'';comment:手机号" json:"phone"`
	Email       string         `gorm:"column:email;uniqueIndex;size:100;not null;comment:邮箱" json:"email"`
	Password    string         `gorm:"column:password;size:128;not null;comment:登录密码" json:"-"`
	Status      bool           `gorm:"column:status;not null;default:true;comment:状态（true=启用，false=禁用）" json:"status"`
	CreatedAt   time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Agent) BeforeCreate(tx *gorm.DB) (err error) {
	a.AgentNo = strings.Join([]string{"MSG", time.Now().Format("20060102"), lo.RandomString(4, lo.NumbersCharset)}, "")
	return nil

}
func (a Agent) TableName() string {
	return "msgbox_agents"
}

func (a *Agent) VerifyPassword(inputPwd string) bool {
	return cryptox.HashCheck(inputPwd, a.Password)
}
