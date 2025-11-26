package models

import (
	"gorm.io/gorm"
	"time"
)

type Template struct {
	ID         int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	AgentID    int64          `gorm:"column:agent_id;uniqueIndex:idx_agent_code;not null;comment:代理商ID" json:"agent_id"`
	ChannelID  int64          `gorm:"column:channel_id;not null;comment:所属通道ID" json:"channel_id"`
	Code       string         `gorm:"column:code;uniqueIndex:idx_agent_code;size:50;not null;comment:模版编码" json:"code"`
	VendorCode string         `gorm:"column:vendor_code;size:100;default:'';comment:厂商模板编码" json:"vendor_code"`
	Signature  string         `gorm:"column:signature;size:64;default:'';comment:签名" json:"signature"`
	Content    string         `gorm:"column:content;type:text;not null;comment:模板内容（含变量占位符）" json:"content"`
	Status     bool           `gorm:"column:status;not null;default:true;comment:是否启用（true=启用，false=禁用）" json:"status"`
	UsedCount  int64          `gorm:"column:used_count;default:0;comment:使用次数" json:"used_count"`
	CreatedAt  time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Channel *Channel `json:"channel,omitempty"`
}

func (t Template) TableName() string {
	return "msgbox_templates"
}
