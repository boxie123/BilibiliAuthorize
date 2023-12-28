package app

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"strings"
)

// NewDevice
//
//	@Description: 生成新的设备信息
//	@param AndroidModel 手机型号
//	@param AndroidBuild 安卓版本
//	@return *Device
//	@return error 错误处理
func NewDevice(AndroidModel string, AndroidBuild string) (*Device, error) {
	device := &Device{
		AndroidModel: AndroidModel,
		AndroidBuild: AndroidBuild,
	}
	device.GenerateFakeBuvid()
	err := device.GetVersions("")
	if err != nil {
		return nil, err
	}
	_, err = device.GetSdkInt()
	if err != nil {
		return nil, err
	}
	return device, nil
}

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
//	@param mod 传参, 若为空则传入 MobiApp
//	@return string build
//	@return string version
//	@return error 错误处理
func (d *Device) GetVersions(mod string) error {
	if mod == "" {
		mod = MobiApp
	}
	client := resty.New()
	resp, err := client.R().
		SetResult(&BilibiliVersionResponse{}).
		SetQueryParam("mobi_app", mod).
		Get("https://app.bilibili.com/x/v2/version")
	if err != nil {
		return err
	}
	result := resp.Result().(*BilibiliVersionResponse)
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
func (d *Device) GetSdkInt() (string, error) {
	AndroidVersion := d.AndroidBuild
	if AndroidVersion == "" {
		return "", fmt.Errorf("请先指定安卓系统版本")
	}
	buildList := strings.Split(AndroidVersion, ".")
	var sdkInt string
	data := SystemSdkIntMap
	for _, li := range buildList {
		value, ok := data[li]
		if !ok {
			if sdkInt != "" {
				return sdkInt, nil
			}
			return "", fmt.Errorf("未找到 %s 的sdk_int", AndroidVersion)
		}
		sdkInt = value["value"].(string)
		data = map[string]map[string]interface{}{
			li: value,
		}
	}
	return sdkInt, nil
}

// BuildUserAgent
//
//	@Description: 构造User-Agent
//	@param device 设备信息
func (d *Device) BuildUserAgent() (string, error) {
	sdkInt, err := d.GetSdkInt()
	if err != nil {
		return "", err
	}
	if d.VersionCode == "" || d.VersionName == "" {
		err = d.GetVersions("")
		if err != nil {
			return "", err
		}
	}
	if d.AndroidModel == "" {
		return "", fmt.Errorf("AndroidModel为空")
	}
	if d.BilibiliBuvid == "" {
		d.GenerateFakeBuvid()
	}
	varMap := map[string]string{
		"ANDROID_BUILD":   d.AndroidBuild,
		"ANDROID_MODEL":   d.AndroidModel,
		"ANDROID_BUILD_M": BuildM,
		"BUVID":           d.BilibiliBuvid,
		"SDK_INT":         sdkInt,
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

	return strTemplate, nil
}
