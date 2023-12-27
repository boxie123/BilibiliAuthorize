package app

import (
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// GenerateFakeBuvid
//
//	@Description: 伪造设备标识码
//	@return string buvid
func (d *Device) GenerateFakeBuvid() {
	randomText := strings.ReplaceAll(uuid.New().String(), "-", "")
	for len(randomText) < 35 {
		randomText += strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	fakeBuvid := "XY" + randomText[:35]
	d.BilibiliBuvid = fakeBuvid
}

// GetVersions
//
//	@Description: 获取最新版本的 build 和 version
//	@param mod 传参, 若为空则传入 "android"
//	@return string build
//	@return string version
//	@return error 错误处理
func (d *Device) GetVersions(mod string) error {
	if mod == "" {
		mod = MobiApp
	}
	client := resty.New()
	resp, err := client.R().
		SetResult(&AppVersionResponse{}).
		SetQueryParam("mobi_app", mod).
		Get("https://app.bilibili.com/x/v2/version")
	if err != nil {
		return err
	}
	result := resp.Result().(*AppVersionResponse)
	if len(result.Data) > 0 {
		d.VersionCode = fmt.Sprintf("%d", result.Data[0].Build)
		d.VersionName = result.Data[0].Version
		return nil
	}

	return fmt.Errorf("no data found")
}

// GetSdkInt
//
//	@Description: 根据安卓系统版本获取 SDK 版本
//	@param AndroidVersion 安卓系统版本
//	@return string sdk版本
//	@return error 错误处理
func (d *Device) GetSdkInt() string {
	AndroidVersion := d.AndroidBuild
	buildList := strings.Split(AndroidVersion, ".")
	var sdkInt string
	data := SystemSdkIntMap
	for _, li := range buildList {
		value, ok := data[li]
		if !ok {
			if sdkInt != "" {
				return sdkInt
			}
			log.Println("未找到" + AndroidVersion + "的sdk_int")
			return ""
		}
		sdkInt = value["value"].(string)
		data = map[string]map[string]interface{}{
			li: value,
		}
	}
	return sdkInt
}

// BuildXBiliAuroraEID
//
//	@Description: 生成 x-bili-aurora-eid
//	@param mid 用户 uid
//	@return string x-bili-aurora-eid
func BuildXBiliAuroraEID(mid string) string {
	length := len(mid)
	byteArr := make([]byte, length)

	if length-1 < 0 {
		return ""
	}

	for i := 0; i < length; i++ {
		s := unicode.ToLower(rune("ad1va46a7lza"[i%12]))
		byteArr[i] = byte(mid[i]) ^ byte(s)
	}

	return base64.StdEncoding.EncodeToString(byteArr)
}

// BuildXBiliTraceID
//
//	@Description: 生成 x-bili-trace-id
//	@param timeStamp
//	@return string
func BuildXBiliTraceID(timeStamp int64) string {
	back6 := strconv.FormatInt(int64(math.Round(float64(timeStamp)/256)), 16)
	front := strings.ReplaceAll(uuid.New().String(), "-", "")
	_data1 := front[6:] + back6[2:]
	_data2 := front[22:] + back6[2:]

	return fmt.Sprintf("%v:%v:0:0", _data1, _data2)
}

// BuildSessionID
//
//	@Description: 构造随机 Session ID
//	@return string SessionID
func BuildSessionID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")[:8]
}

// BuildUserAgent
//
//	@Description: 构造User-Agent
//	@param device 设备信息
func (d *Device) BuildUserAgent() string {
	varMap := map[string]string{
		"ANDROID_BUILD":   d.AndroidBuild,
		"ANDROID_MODEL":   d.AndroidModel,
		"ANDROID_BUILD_M": BuildM,
		"BUVID":           d.BilibiliBuvid,
		"SDK_INT":         d.GetSdkInt(),
		"VERSION_CODE":    d.VersionCode,
		"CHANNEL":         Channel,
		"SESSION_ID":      BuildSessionID(),
		"VERSION_NAME":    d.VersionName,
	}
	// 定义待替换的字符串模板
	strTemplate := UserAgentFormat
	// 执行替换
	for k, v := range varMap {
		strTemplate = strings.ReplaceAll(strTemplate, "{"+k+"}", v)
	}

	return strTemplate
}
