package web

import (
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
	param := map[string]interface{}{
		"foo": "114",
		"bar": "514",
		"zab": 1919810,
		"wts": 1702204169,
	}
	cache.Store("imgKey", "7cd084941338484aae1ad9425b84077c")
	cache.Store("subKey", "4932caff0ff746eab6f01bf08b70ac45")
	lastUpdateTime = time.Now()
	param = ParamSign(param)
	if param["w_rid"] != "8f6f2b5b3d485fe1886cec6a0be8c5d4" {
		t.Errorf("Wbi签名错误: w_rid = %s", param["w_rid"])
	}
}
