package sms

import (
	"fmt"
	"time"
)

const (
	// SMS_TYEP_JD_CLOUD 京东云短信
	SMS_TYEP_JD_CLOUD = "JDCloud SMS"
)

type SendMessageResp struct {
	Content  string // 发送内容
	SplitNum int    // 拆分条数
	UniqueId string // 唯一id
	Code     string // 错误码
	Message  string // 错误消息
}

type PullReceiptResp struct {
	Status  string
	Message string
	Items   []*PullItem
}

type PullItem struct {
	UniqueId    string    // 唯一id
	Code        string    // 错误码
	SendTime    time.Time // 发送时间
	ReceiveTime time.Time // 接收时间
}

type Client interface {
	SendMessage(param map[string]string, targetPhoneNumber ...string) (*SendMessageResp, error) // 发送消息
	PullReceipt() (*PullReceiptResp, error)                                                     // 拉取回执
}

// NewSmsClient 创建短信客户端 (兼容原先的短信客户端)
func NewSmsClient(provider string, accessId string, accessKey string, sign string, template string, other ...string) (Client, error) {
	switch provider {
	case SMS_TYEP_JD_CLOUD:
		return NewJDCloudClient(accessId, accessKey, sign, template, other)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
		//return go_sms_sender.NewSmsClient(provider, accessId, accessKey, sign, template, other...)
	}
}
