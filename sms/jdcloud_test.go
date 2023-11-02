package sms

import (
	"testing"
)

func testNewJdClient() (*JDCloudClient, error) {
	var (
		accessId  = ""
		accessKey = ""
		sign      = ""
		template  = ""
		apiAddr   = ""
	)

	return NewJDCloudClient(accessId, accessKey, sign, template, []string{apiAddr})
}

func TestJDCloudClient_SendMessage(t *testing.T) {
	var (
		phones = []string{"13788992526"}
		params = map[string]string{
			"code":    "6666",
			"min_num": "test",
		}
	)

	c, err := testNewJdClient()
	if err != nil {
		t.Errorf("new client err = %v", err)
	}
	if _, err = c.SendMessage(params, phones...); err != nil {
		t.Errorf("SendMessage() error = %v", err)
	}
}

// 231027105529108511200015
// pull receipt resp: sms.jdPullResp{Status:"00", Cause:"成功", Size:1, Result:[]sms.jdPullItem{sms.jdPullItem{ReqId:"231027105529108511200015", Phone:"13788992526", State:"1", Code:"JD:0015", SendTime:"2023-10-27 10:55:29", ReceiveTime:"2023-10-27 10:55:29", OrderId:"1"}}}--- PASS: TestJDCloudClient_PullReceipt (0.12s)
func TestJDCloudClient_PullReceipt(t *testing.T) {
	c, err := testNewJdClient()
	if err != nil {
		t.Errorf("new client err = %v", err)
	}
	if _, err = c.PullReceipt(); err != nil {
		t.Errorf("PullReceipt() error = %v", err)
	}
}
