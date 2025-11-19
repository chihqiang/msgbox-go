package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	ContentTypeTxt = 1 // 纯文本（默认）
	ContentTypeMd  = 2 // Markdown 格式
)

// 内容类型名称映射（便于方法中快速查找）
var contentTypeNameMap = map[int]string{
	ContentTypeTxt: "纯文本",
	ContentTypeMd:  "Markdown",
}

type Template struct {
	ID          int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	AgentID     int64          `gorm:"column:agent_id;uniqueIndex:idx_agent_code;not null;comment:代理商ID" json:"agent_id"`
	Code        string         `gorm:"column:code;uniqueIndex:idx_agent_code;size:50;not null;comment:模版编码" json:"code"`
	VendorCode  string         `gorm:"column:vendor_code;size:100;default:'';comment:厂商模板编码" json:"vendor_code"`
	Signature   string         `gorm:"column:signature;size:64;default:'';comment:签名" json:"signature"`
	Content     string         `gorm:"column:content;type:text;not null;comment:模板内容（含变量占位符）" json:"content"`
	ContentType int            `gorm:"column:content_type;type:tinyint(1);default:1;comment:发送消息的类型(1=纯文本,2=markdown)" json:"content_type"`
	Status      bool           `gorm:"column:status;not null;default:true;comment:是否启用（true=启用，false=禁用）" json:"status"`
	UsedCount   int64          `gorm:"column:used_count;default:0;comment:使用次数" json:"used_count"`
	CreatedAt   time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Channels []*Channel `gorm:"many2many:msgbox_channel_templates;" json:"channels"`
}

func (t Template) TableName() string {
	return "msgbox_templates"
}

// ContentTypeTitle 根据 ContentType 获取对应的中文名称
// 返回值：类型名称（默认返回"未知类型"）
func (t *Template) ContentTypeTitle() string {
	name, ok := contentTypeNameMap[t.ContentType]
	if ok {
		return name
	}
	return "未知类型"
}
