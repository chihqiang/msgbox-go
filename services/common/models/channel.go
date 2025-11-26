package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Channel struct {
	ID         int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	AgentID    int64          `gorm:"column:agent_id;uniqueIndex:idx_agent_code;not null;comment:代理商ID" json:"agent_id"`
	Code       string         `gorm:"column:code;uniqueIndex:idx_agent_code;size:50;not null;comment:通道编码" json:"code"`
	Name       string         `gorm:"column:name;size:50;default:'';comment:通道名称" json:"name"`
	VendorName string         `gorm:"column:vendor_name;size:50;not null;comment:服务商名称" json:"vendor_name"`
	Config     datatypes.JSON `gorm:"column:config;type:JSON;not null;comment:通道配置" json:"config"`
	Status     bool           `gorm:"column:status;not null;default:true;comment:状态（true=启用，false=禁用）" json:"status"`
	CreatedAt  time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	Agent     *Agent      `gorm:"foreignKey:AgentID" json:"agent,omitempty"`
	Templates []*Template `gorm:"foreignKey:ChannelID" json:"templates"`
}

func (a Channel) TableName() string {
	return "msgbox_channels"
}
