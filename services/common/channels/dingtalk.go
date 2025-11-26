package channels

import (
	"chihqiang/msgbox-go/pkg/clientx"
	"chihqiang/msgbox-go/pkg/htmlx"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type DingTalkSender struct {
	Endpoint    string `json:"endpoint" ui:"label=Webhook地址;type=text;required;placeholder=请输入 Webhook;default=https://oapi.dingtalk.com/robot/send"`
	AccessToken string `json:"access_token" ui:"label=AccessToken;type=text;required;placeholder=请输入 AccessToken"`
	Secret      string `json:"secret" ui:"label=Secret;type=text;placeholder=请输入 Secret（可选）"`
}

func (d *DingTalkSender) GetName() string {
	return "dingtalk"
}
func (d *DingTalkSender) GetLabel() string {
	return "钉钉机器人"
}

func (d *DingTalkSender) SetConfig(config map[string]any) error {
	return htmlx.MapSet(d, config)
}

func (d *DingTalkSender) url() *url.URL {
	uri, _ := url.Parse(d.Endpoint)
	query := url.Values{}
	query.Set("access_token", d.AccessToken)
	if d.Secret != "" {
		milliTimestamp := time.Now().UnixNano() / 1e6
		stringToSign := fmt.Sprintf("%s\n%s", strconv.Itoa(int(milliTimestamp)), d.Secret)
		mac := hmac.New(sha256.New, []byte(d.Secret))
		mac.Write([]byte(stringToSign))
		sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
		query.Set("timestamp", strconv.Itoa(int(milliTimestamp)))
		query.Set("sign", sign)
	}
	uri.RawQuery = query.Encode()
	return uri
}

func (d *DingTalkSender) Send(message IMessage) (resp map[string]any, err error) {
	req := map[string]any{
		"msgtype": "text",
		"text": map[string]any{
			"content": message.GetContent(),
		},
	}
	receiver := message.GetReceiver()
	switch receiver {
	case "all":
		req["isAtAll"] = true
	default:
		req["atMobiles"] = []string{receiver}
	}
	response, err := clientx.PostJSON(context.Background(), d.url().String(), req)
	if err != nil {
		return resp, err
	}
	err = json.NewDecoder(response.Body).Decode(&resp)
	if errCode, ok := resp["errcode"].(float64); ok && errCode != 0 {
		err = errors.New(resp["errmsg"].(string))
	}
	return resp, err
}
