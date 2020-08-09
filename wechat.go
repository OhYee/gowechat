package wechat

import (
	"crypto/sha1"
	"sort"
	"strings"
)

type Wechat struct {
	Token  string
	AESKey string
}

// CheckSignature 检查签名是否合法
//
// 如果验证成功，应该原样返回接收到的 echostr
// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html
func (wc *Wechat) CheckSignature(signature string, timestamp string, nonce string) bool {
	array := []string{wc.Token, timestamp, nonce}

	sort.SliceStable(array, func(i, j int) bool {
		return array[i] < array[j]
	})

	hash := sha1.New()
	hash.Write([]byte(strings.Join(array, "")))

	return string(hash.Sum([]byte{})) == signature
}
