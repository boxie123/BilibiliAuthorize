package app

import (
	"fmt"
	"testing"
)

func TestParamSign(t *testing.T) {
	param := map[string]interface{}{
		"id":   114514,
		"str":  "1919810",
		"test": "いいよ，こいよ",
	}
	param = ParamSign(param)
	fmt.Println(param)
	if param["sign"] != "01479cf20504d865519ac50f33ba3a7d" {
		t.Error("APP签名错误")
	}
}

func TestGetVersions(t *testing.T) {
	build, _, err := GetVersions("")
	if err != nil {
		t.Error(err)
	}
	if build != "7580300" {
		t.Errorf("获取版本号失败: build = %s", build)
	}
}
