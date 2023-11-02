package sms

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"
)

func post(url string, param []byte, headers map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(param))
	if err != nil {
		return nil, err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateContent 生成短信内容
func GenerateContent(sign, template string, params map[string]string) string {
	return fmt.Sprintf("【%s】%s", sign, FillInTemplate(template, params))
}

// FillInTemplate 模版填充
func FillInTemplate(template string, params map[string]string) string {
	return os.Expand(template, func(k string) string { return params[k] })
}

// GetCurrentDate 格式:yyyymmddhhmmss
func GetCurrentDate() string {
	return time.Now().Format("20060102150405")
}

// SplitNumContent 拆分短信内容
func SplitNumContent(content string) (int, error) {
	contLen := len([]rune(content))

	if contLen > MAX_SMS_MESSAGE_LENGTH {
		return 0, fmt.Errorf("sms message is too long")
	}

	// 普通短信: 内容为 70 字符(含)以内，按 1 条计费;
	// 长短信计费规则:短信内容 70 字符以上，1000 字符(含)以内，每 67 字 符 1 条进行计费
	var splitNum int
	if contLen <= 70 {
		splitNum = 1
	} else {
		splitNum = int(math.Ceil(float64(contLen) / float64(67)))
	}
	return splitNum, nil
}

// DateString2Time 日期字符串转时间
// 2023-10-27 10:55:29
func DateString2Time(dateStr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
