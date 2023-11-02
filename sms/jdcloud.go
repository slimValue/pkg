package sms

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type JDCloudClient struct {
	accessId   string
	accessKey  string
	sign       string
	template   string
	apiAddress string
}

type jdSmsReq struct {
	Name    string // 账号
	Pwd     string // 密码 MD5(pwd+mttime)
	Phone   string
	Context string
	Mttime  string // 提交时间  yyyyMMddhhmmss
	Rpttype string // 1:json
}

type jdSmsResp struct {
	ReqCode string `json:"reqCode"`
	ReqMsg  string `json:"reqMsg"`
	ReqId   string `json:"reqId"` // 批次id
}

const (
	RPTTYPE_JSON           = "1"
	HEADER_FORM_URLENCODED = "application/x-www-form-urlencoded"

	SMS_REQ_Uri  = "HttpSmsMt" // 发送短信
	SMS_PULL_Uri = "rptpull"   // 接收回执

	MAX_SMS_MESSAGE_LENGTH = 1000 // 最长sms消息长度
)

type jdPullReq struct {
	Account   string // 账号
	Token     string // 请求令牌 MD5HEX (密码+timestamp)
	Timestamp string // 时间戳 yyyyMMddhhmmss
}

type jdPullItem struct {
	ReqId       string
	Phone       string
	State       string
	Code        string
	SendTime    string
	ReceiveTime string
	OrderId     string
}

type jdPullResp struct {
	Status string
	Cause  string
	Size   int
	Result []jdPullItem
}

func NewJDCloudClient(accessId string, accessKey string, sign string, template string, other []string) (*JDCloudClient, error) {
	if len(other) < 1 {
		return nil, fmt.Errorf("missing parameter: apiAddress")
	}

	apiAddress := strings.TrimRight(other[0], "/")

	jdClient := &JDCloudClient{
		accessId:   accessId,
		accessKey:  accessKey,
		sign:       sign,
		template:   template,
		apiAddress: apiAddress,
	}

	return jdClient, nil
}

func (c *JDCloudClient) SendMessage(param map[string]string, targetPhoneNumber ...string) (*SendMessageResp, error) {
	reqUrl := fmt.Sprintf("%s/%s", c.apiAddress, SMS_REQ_Uri)

	headers := map[string]string{
		"Content-Type": HEADER_FORM_URLENCODED,
	}

	// 国内手机号 最多一次100个
	if len(targetPhoneNumber) > 100 {
		return nil, fmt.Errorf("too many phone numbers")
	}

	var phones []string
	for _, phone := range targetPhoneNumber {
		if strings.HasPrefix(phone, "+86") {
			phone = phone[3:]
		} else if strings.HasPrefix(phone, "+") { // 暂不支持国际区号
			return nil, fmt.Errorf("unsupported country code")
		}
		phones = append(phones, phone)
	}

	curTime := GetCurrentDate()
	pwd := GetMd5String(c.accessKey + curTime)

	// post 参数见 jdSmsReq
	body := url.Values{}
	body.Add("pwd", pwd)
	body.Add("mttime", curTime)
	body.Add("name", c.accessId)
	body.Add("rpttype", RPTTYPE_JSON)
	body.Add("phone", strings.Join(phones, ","))
	bodyStr := body.Encode()

	// content 不需要URL编码
	content := GenerateContent(c.sign, c.template, param)

	splitNum, err := SplitNumContent(content)

	//content = url.QueryEscape(content)
	bodyStr += fmt.Sprintf("&%s=%s", "content", content)
	bodyBytes := []byte(bodyStr)

	data, err := post(reqUrl, bodyBytes, headers)
	if err != nil {
		return nil, err
	}

	var resp jdSmsResp
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	// {"ReqMsg":"提交参数异常，手机号参数不正确","ReqId":"231026120412108511000013","ReqCode":"02"}
	fmt.Printf("send message resp: %#v", resp)

	return &SendMessageResp{
		Content:  content,
		SplitNum: splitNum,
		UniqueId: resp.ReqId,
		Code:     resp.ReqCode,
		Message:  resp.ReqMsg,
	}, nil
}

// PullReceipt 拉取回执
func (c *JDCloudClient) PullReceipt() (*PullReceiptResp, error) {
	reqUrl := fmt.Sprintf("%s/%s", c.apiAddress, SMS_PULL_Uri)

	headers := map[string]string{
		"Content-Type": HEADER_FORM_URLENCODED,
	}

	// 格式:yyyymmddhhmmss
	curTime := GetCurrentDate()

	body := url.Values{}
	body.Add("account", c.accessId)
	body.Add("timestamp", curTime)
	body.Add("token", GetMd5String(c.accessKey+curTime))
	bodyBytes := []byte(body.Encode())

	data, err := post(reqUrl, bodyBytes, headers)
	if err != nil {
		return nil, err
	}

	var resp jdPullResp
	if err = json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	var items []*PullItem
	if len(resp.Result) >= 0 {
		for _, v := range resp.Result {
			st, _ := DateString2Time(v.SendTime)
			rt, _ := DateString2Time(v.ReceiveTime)

			items = append(items, &PullItem{
				UniqueId:    v.ReqId,
				Code:        v.Code,
				SendTime:    st,
				ReceiveTime: rt,
			})
		}

	}

	fmt.Printf("pull receipt resp: %#v", resp)

	return &PullReceiptResp{
		Status:  resp.Status,
		Message: resp.Cause,
		Items:   items,
	}, nil
}
