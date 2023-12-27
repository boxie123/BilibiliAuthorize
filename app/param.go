package app

import (
	"github.com/boxie123/BilibiliAuthorize/utils"
)

const (
	AppKey      = "1d8b6e7d45233436"
	AppSec      = "560c52ccd288fed045859ed18bffd973"
	Platform    = "android"
	NeuronAppId = 1
	MobiApp     = "android"
)

// ParamSign 为参数添加 appkey 和 sign
func ParamSign(param map[string]interface{}) map[string]interface{} {
	_, ok := param["appkey"]
	if !ok {
		param["appkey"] = AppKey
	}
	paramQuery := utils.ParamEncode(param)
	sign := utils.Md5(paramQuery + AppSec)
	param["sign"] = sign
	return param
}
