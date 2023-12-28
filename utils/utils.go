package utils

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

// ParamEncode 将 map[string]string 类型的参数按 key 排序, 编码为 url encoded 形式
func ParamEncode(param map[string]string) string {
	v := url.Values{}
	for key, value := range param {
		v.Add(key, value)
	}
	s := v.Encode()
	// 将空格用 "%20" 编码, 而不是 "+"
	s = strings.Replace(s, "+", "%20", -1)
	return s
}

// Md5 md5加密
func Md5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return md5str
}
