package sign

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type Signer struct {
	body      url.Values // 签名参数体
	secretKey string     // 签名密钥
}

func NewSigner() *Signer {
	return &Signer{
		body: make(url.Values),
	}
}

// SetTimestamp 设置时间戳参数
func (s *Signer) SetTimestamp(ts int64) *Signer {
	return s.AddBody(KeyNameTimeStamp, strconv.FormatInt(ts, 10))
}

// SetApp 设置应用
func (s *Signer) SetApp(app string) *Signer {
	return s.AddBody(KeyNameApp, app)
}

// SetOwner 设置租户
func (s *Signer) SetOwner(owner string) *Signer {
	return s.AddBody(KeyNameOwner, owner)
}

// SetSecret 设置签名密钥
func (s *Signer) SetSecret(appSecret string) *Signer {
	s.secretKey = appSecret
	return s
}

// SetBody 设置整个参数体Body对象。
func (s *Signer) SetBody(body url.Values) {
	for k, v := range body {
		s.body[k] = v
	}
}

// GetBody 返回Body内容
func (s *Signer) GetBody() url.Values {
	return s.body
}

func (s *Signer) AddBody(key string, value string) *Signer {
	s.body[key] = []string{value}
	return s
}

func (s *Signer) MakeSignedQuery() string {
	body := s.getSortedBodyString()
	sign := s.GetSignature()
	return body + "&" + KeyNameSign + "=" + sign
}

func (s *Signer) GetRawSignature() []byte {
	return Md5Sign(s.getSortedBodyString())
}

func (s *Signer) GetSignature() string {
	sign := fmt.Sprintf("%x",s.GetRawSignature())
	return sign
}

func (s *Signer) getSortedBodyString() string {
	return SortKVPairs(s.body)
}

func (s *Signer) GetSignedQuery() string {
	return s.MakeSignedQuery()
}

func (s *Signer) GetSignBodyString() string {
	return s.getSortedBodyString()
}

func SortKVPairs(m url.Values) string {
	size := len(m)
	if size == 0 {
		return ""
	}
	keys := make([]string, size)
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	pairs := make([]string, size)
	for i, key := range keys {
		pairs[i] = key + "=" + strings.Join(m[key], ",")
	}
	return strings.Join(pairs, "&")
}

func Md5Sign(body string) []byte {
	m := md5.New()
	m.Write([]byte(body))
	return m.Sum(nil)
}
