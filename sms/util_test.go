package sms

import (
	"fmt"
	"testing"
)

func TestGetMd5String(t *testing.T) {
	var rsp = &jdSmsResp{
		ReqCode: "1",
	}
	fmt.Printf("%#v", rsp)
	if got := GetMd5String("test"); got != "098f6bcd4621d373cade4e832627b4f6" {
		t.Log(got)
		t.Errorf("HashMd5() = %x", got)
	}
}

func TestFillInTemplate(t *testing.T) {
	tpl := "${code}为您的登录验证码，请于${min_num}分钟内使用，如非本人操作，请忽略本短信。回复BK或在APP设置中退订"
	params := map[string]string{
		"code":    "1234",
		"min_num": "5",
	}
	got := FillInTemplate(tpl, params)
	t.Logf("FillTemplate() = %v", got)
}

func TestSplitNumContent(t *testing.T) {
	content := "【京东云】您的验证码是1234，有效期5分钟，如非本人操作，请忽略本短信。回复BK或在APP设置中退订"
	got, err := SplitNumContent(content)
	if err != nil {
		t.Errorf("SplitNumContent() error = %v", err)
		return
	}

	t.Logf("SplitNumContent() = %v", got)
}

func TestDateString2Time(t *testing.T) {
	got, err := DateString2Time("2023-10-27 10:55:29")
	if err != nil {
		t.Errorf("TestDateString2Time() error = %v", err)
		return
	}
	t.Logf("TestDateString2Time() = %v", got)
}