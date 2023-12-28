package web

import (
	"github.com/go-resty/resty/v2"
	"testing"
	"time"
)

func TestGetMixinKey(t *testing.T) {
	ImgKey := "7cd084941338484aae1ad9425b84077c"
	SubKey := "4932caff0ff746eab6f01bf08b70ac45"
	mixinKey := getMixinKey(ImgKey + SubKey)
	if mixinKey != "ea1db124af3c7062474693fa704f4ff8" {
		t.Errorf("Wbi签名错误: mixinKey = %s", mixinKey)
	}
}

func TestParamSign(t *testing.T) {
	param := map[string]string{
		"foo": "114",
		"bar": "514",
		"zab": "1919810",
		"wts": "1702204169",
	}
	cache.Store("imgKey", "7cd084941338484aae1ad9425b84077c")
	cache.Store("subKey", "4932caff0ff746eab6f01bf08b70ac45")
	lastUpdateTime = time.Now()
	param = ParamSign(param)
	if param["w_rid"] != "8f6f2b5b3d485fe1886cec6a0be8c5d4" {
		t.Errorf("Wbi签名错误: w_rid = %s", param["w_rid"])
	}
}

func TestRequest(t *testing.T) {
	type Result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	result := Result{}
	param := map[string]string{
		"vmid":         "1485569",
		"web_location": "333.999",
	}
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Origin":     "https://space.bilibili.com",
		"Referer":    "https://space.bilibili.com/1485569/",
	}
	param = ParamSign(param)
	client := resty.New()
	_, err := client.R().
		SetResult(&result).
		SetQueryParams(param).
		SetHeaders(headers).
		Get("https://api.bilibili.com/x/relation/stat")
	if err != nil {
		t.Error(err)
	}
	if result.Code != 0 {
		t.Errorf("Code: %d, Message: %s", result.Code, result.Message)
	}
}
