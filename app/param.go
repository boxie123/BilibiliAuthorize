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

// ParamSign
//
//	@Description: 向参数映射中添加 appkey 和 sign 鉴权签名
//	@param param 需要加签名的参数映射
//	@return map[string]interface{}
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
