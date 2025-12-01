package senders

import (
	"chihqiang/msgbox-go/pkg/clientx"
	"chihqiang/msgbox-go/pkg/htmlx"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type WorkWxSender struct {
	URL string `json:"url" ui:"label=Webhook地址;type=text;required;placeholder=请输入接口请求地址;default=https://qyapi.weixin.qq.com/cgi-bin/webhook/send"`
	Key string `json:"key" ui:"label=key;type=text;required;placeholder=调用接口凭证"`
}

func (w *WorkWxSender) GetName() string {
	return "workwx"
}

func (w *WorkWxSender) GetLabel() string {
	return "企业微信机器人"
}

func (w *WorkWxSender) SetConfig(config map[string]any) error {
	return htmlx.MapSet(w, config)
}

// buildWebhookURL 构建完整的 Webhook URL
func (w *WorkWxSender) buildWebhookURL() string {
	webhook := strings.TrimSpace(w.URL)
	key := strings.TrimSpace(w.Key)
	// 处理 URL 中已包含 key 参数的情况
	if strings.Contains(webhook, "key=") {
		return webhook
	}
	// 拼接 key 参数（处理 URL 是否以 ? 结尾）
	if strings.HasSuffix(webhook, "?") {
		return fmt.Sprintf("%skey=%s", webhook, key)
	} else if strings.Contains(webhook, "?") {
		return fmt.Sprintf("%s&key=%s", webhook, key)
	} else {
		return fmt.Sprintf("%s?key=%s", webhook, key)
	}
}
func (w *WorkWxSender) Send(message IMessage) (resp map[string]any, err error) {
	req := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content":               message.GetContent(),
			"mentioned_mobile_list": []string{message.GetReceiver()},
		},
	}
	response, err := clientx.PostJSON(context.Background(), w.buildWebhookURL(), req)
	if err != nil {
		return resp, err
	}
	_ = json.NewDecoder(response.Body).Decode(&resp)
	if errCode, ok := resp["errcode"].(float64); ok && errCode != 0 {
		err = errors.New(resp["errmsg"].(string))
	}
	return resp, err
}
