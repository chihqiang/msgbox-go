package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	SendBatchStatusPending  = 1 // 待发送
	SendBatchStatusSending  = 2 // 发送中
	SendBatchStatusFinished = 3 // 已完成
	SendBatchStatusFailed   = 4 // 异常/失败
)

const (
	SendRecordStatusPending = 1 // 待发送
	SendRecordStatusSending = 2 // 发送中
	SendRecordStatusSuccess = 3 // 成功
	SendRecordStatusFailed  = 4 // 失败
)

type SendBatch struct {
	ID            int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	AgentID       int64          `gorm:"column:agent_id;not null;index;comment:代理商ID" json:"agent_id"`
	ChannelID     int64          `gorm:"column:channel_id;not null;index;comment:通道ID" json:"channel_id"`
	TemplateID    int64          `gorm:"column:template_id;comment:模板ID" json:"template_id"`
	BatchNo       string         `gorm:"column:batch_no;size:64;uniqueIndex;not null;comment:批次唯一编号" json:"batch_no"`
	TraceID       string         `gorm:"column:trace_id;size:100;not null;comment:链路ID" json:"trace_id"`
	TotalCount    int            `gorm:"column:total_count;default:0;comment:总消息条数" json:"total_count"`
	SuccessCount  int            `gorm:"column:success_count;default:0;comment:发送成功条数" json:"success_count"`
	FailCount     int            `gorm:"column:fail_count;default:0;comment:发送失败条数" json:"fail_count"`
	Status        int            `gorm:"column:status;not null;default:1;comment:批次状态(1=待发送,2=发送中,3=完成,4=异常)" json:"status"`
	ScheduledTime *time.Time     `gorm:"column:scheduled_time;comment:计划发送时间" json:"scheduled_time"`
	SendStartTime *time.Time     `gorm:"column:send_start_time;comment:实际开始发送时间" json:"send_start_time"`
	SendEndTime   *time.Time     `gorm:"column:send_end_time;comment:实际结束发送时间" json:"send_end_time"`
	CreatedAt     time.Time      `gorm:"column:created_at;autoCreateTime:nano" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;autoUpdateTime:nano" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`

	Agent    *Agent        `gorm:"foreignKey:AgentID" json:"agent,omitempty"`
	Channel  *Channel      `gorm:"foreignKey:ChannelID" json:"channel,omitempty"`
	Template *Template     `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
	Records  []*SendRecord `gorm:"foreignKey:BatchID" json:"records,omitempty"`
}

func (b *SendBatch) TableName() string {
	return "msgbox_send_batches"
}

type SendRecord struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	BatchID       int64          `gorm:"column:batch_id;index;comment:所属批次ID" json:"batch_id"`
	AgentID       int64          `gorm:"column:agent_id;not null;index;comment:代理商ID" json:"agent_id"`
	ChannelID     int64          `gorm:"column:channel_id;not null;index;comment:通道ID" json:"channel_id"`
	TemplateID    int64          `gorm:"column:template_id;not null;comment:模板ID，可空" json:"template_id"`
	TraceID       string         `gorm:"column:trace_id;size:100;not null;comment:链路ID" json:"trace_id"`
	Receiver      string         `gorm:"column:receiver;size:100;not null;comment:发送目标（手机号/邮箱）" json:"receiver"`
	VendorName    string         `gorm:"column:vendor_name;size:50;not null;comment:服务商名称" json:"vendor_name"`
	ChannelConfig datatypes.JSON `gorm:"column:channel_config;type:JSON;not null;comment:通道配置" json:"channel_config"`
	VendorCode    string         `gorm:"column:vendor_code;size:100;default:'';comment:厂商模板编码" json:"vendor_code"`
	Signature     string         `gorm:"column:signature;size:64;default:'';comment:签名" json:"signature"`
	Title         string         `gorm:"column:title;size:255;default:'';comment:消息标题" json:"title"`
	Content       string         `gorm:"column:content;type:text;not null;comment:最终发送内容" json:"content"`
	Variables     datatypes.JSON `gorm:"column:variables;type:json;comment:模板渲染参数" json:"variables"`
	Extra         datatypes.JSON `gorm:"column:extra;type:json;comment:扩展参数" json:"extra"`
	Status        int            `gorm:"column:status;not null;default:1;comment:消息状态(1=待发送,2=发送中,3=成功,4=失败)" json:"status"`
	SendTime      *time.Time     `gorm:"column:send_time;comment:发送动作时间" json:"send_time"`
	Error         string         `gorm:"column:error;size:255;default:'';comment:错误内容" json:"error"`
	Response      datatypes.JSON `gorm:"column:response;type:json;comment:服务商原始响应" json:"response"`
	DeliveryTime  *time.Time     `gorm:"column:delivery_time;comment:回执回调时间" json:"delivery_time"`
	DeliveryRaw   datatypes.JSON `gorm:"column:delivery_raw;type:json;comment:回执原始内容" json:"delivery_raw"`
	CreatedAt     time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Batch         *SendBatch     `gorm:"foreignKey:BatchID" json:"batch,omitempty"`
	Agent         *Agent         `gorm:"foreignKey:AgentID" json:"agent,omitempty"`
	Channel       *Channel       `gorm:"foreignKey:ChannelID" json:"channel,omitempty"`
	Template      *Template      `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
}

func (sr *SendRecord) StatusMsg() string {
	switch sr.Status {
	case SendRecordStatusSending:
		return "发送中"
	case SendRecordStatusSuccess:
		return "成功"
	case SendRecordStatusFailed:
		return "失败"
	default:
		return "待发送"
	}
}

func (sr *SendRecord) GetReceiver() string {
	return sr.Receiver
}

func (sr *SendRecord) GetSignature() string {
	return sr.Signature
}

func (sr *SendRecord) GetVendorCode() string {
	return sr.VendorCode
}

func (sr *SendRecord) GetTitle() string {
	return sr.Title
}

func (sr *SendRecord) GetContent() string {
	return sr.Content
}

func (sr *SendRecord) GetVariables() map[string]any {
	return DataTypesToMap(sr.Variables)
}

func (sr *SendRecord) GetExtra() map[string]any {
	return DataTypesToMap(sr.Extra)
}

func (sr *SendRecord) TableName() string {
	return "msgbox_send_records"
}
