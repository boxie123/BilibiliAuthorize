package app

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"testing"
	"time"
)

func TestParamSign(t *testing.T) {
	param := map[string]string{
		"id":   "114514",
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
	device := &Device{}
	err := device.GetVersions("")
	if err != nil {
		t.Error(err)
	}
	if device.VersionCode != "7580300" {
		t.Errorf("获取版本号失败: build = %s", device.VersionCode)
	}
}

func TestDevice(t *testing.T) {
	device, err := NewDevice("JKM-AL00", "9")
	if err != nil {
		t.Error(err)
		return
	}
	device.GenerateFakeBuvid()
	err = device.GetVersions("")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(device)
}

func TestDevice_BuildUserAgent(t *testing.T) {
	device, err := NewDevice("JKM-AL00", "9")
	if err != nil {
		t.Error(err)
		return
	}
	ua, err := device.BuildUserAgent()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(ua)
}

func TestRequest(t *testing.T) {
	type Result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	device, err := NewDevice("JKM-AL00", "9")
	if err != nil {
		t.Error(err)
		return
	}
	client := resty.New()
	result := Result{}
	timeStamp := time.Now().Unix()
	params := map[string]string{
		"actionKey":    "appkey",
		"build":        device.VersionCode,
		"c_locale":     "zh-Hans_CN",
		"device":       MobiApp,
		"device_model": device.AndroidModel,
		"disable_rcmd": "0",
		"jump_from":    "24001",
		"mobi_app":     MobiApp,
		"platform":     Platform,
		"room_id":      "1184275",
		"s_locale":     "zh-Hans_CN",
		"statistics":   fmt.Sprintf(`{"appId":1,"version":"%s","abtest":"","platform":1}`, device.VersionName),
		"ts":           strconv.FormatInt(timeStamp, 10),
	}
	params = ParamSign(params)
	ua, err := device.BuildUserAgent()
	if err != nil {
		t.Error(err)
		return
	}
	headers := map[string]string{
		"Host":            "api.live.bilibili.com",
		"Connection":      "keep-alive",
		"Content-Type":    "application/x-www-form-urlencoded",
		"APP-KEY":         MobiApp,
		"Buvid":           device.BilibiliBuvid,
		"User-Agent":      ua,
		"ENV":             "prod",
		"Session_ID":      BuildSessionID(),
		"x-bili-trace-id": BuildXBiliTraceID(timeStamp),
		"Accept-Encoding": "gzip, deflate, br",
	}
	_, err = client.R().
		SetResult(&result).
		SetQueryParams(params).
		SetHeaders(headers).
		Get("https://api.live.bilibili.com/xlive/app-room/v1/index/getInfoByRoom")
	if err != nil {
		t.Error(err)
	}
	if result.Code != 0 {
		t.Errorf("Code: %d, Message: %s", result.Code, result.Message)
	}
}
