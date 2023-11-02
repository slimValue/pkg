package sign

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Verifier struct {
	body    url.Values
	timeout time.Duration // 签名过期时间
}

func NewVerifier() *Verifier {
	return &Verifier{
		body:    make(url.Values),
		timeout: time.Minute * 5,
	}
}

// ParseQuery 将参数字符串解析成参数列表
func (v *Verifier) ParseQuery(requestUri string) error {
	requestQuery := ""
	idx := strings.Index(requestUri, "?")
	if idx > 0 {
		requestQuery = requestUri[idx+1:]
	}
	query, err := url.ParseQuery(requestQuery)
	if nil != err {
		return err
	}
	v.ParseValues(query)
	return nil
}

// ParseValues 将Values参数列表解析成参数Map。如果参数是多值的，则将它们以逗号Join成字符串
func (v *Verifier) ParseValues(values url.Values) {
	for key, value := range values {
		v.body[key] = value
	}
}

func (v *Verifier) GetApp() string {
	return v.MustString(KeyNameApp)
}

func (v *Verifier) GetOwner() string {
	return v.MustString(KeyNameOwner)
}


func (v *Verifier) GetSign() string {
	return v.MustString(KeyNameSign)
}

func (v *Verifier) GetTimestamp() int64 {
	return v.MustInt64(KeyNameTimeStamp)
}

// SetTimeout 设置签名校验过期时间
func (v *Verifier) SetTimeout(timeout time.Duration) *Verifier {
	v.timeout = timeout
	return v
}

func (v *Verifier) CheckTimeStamp() error {
	timestamp := v.GetTimestamp()
	thatTime := time.Unix(timestamp, 0)
	if time.Now().Sub(thatTime) > v.timeout {
		return errors.New(fmt.Sprintf("TIMESTAMP_TIMEOUT:<%d>", timestamp))
	}
	return nil
}

func (v *Verifier) MustString(key string) string {
	if ss := v.MustStrings(key); len(ss) == 0 {
		return ""
	} else {
		return ss[0]
	}
}

func (v *Verifier) MustStrings(key string) []string {
	return v.body[key]
}

// MustInt64 获取Int64值
func (v *Verifier) MustInt64(key string) int64 {
	val := v.MustString(key)
	sv := fmt.Sprintf("%v", val)
	if iv, err := strconv.ParseInt(sv, 10, 64); nil != err {
		return 0
	} else {
		return iv
	}
}

func (v *Verifier) GetBodyWithoutSign() url.Values {
	out := make(url.Values)
	for k, v := range v.body {
		if k != KeyNameSign {
			out[k] = v
		}
	}
	return out
}

func (v *Verifier) GetBody() url.Values {
	out := make(url.Values)
	for k, v := range v.body {
		out[k] = v
	}
	return out
}