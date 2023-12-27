package app

var SystemSdkIntMap = map[string]map[string]interface{}{
	"13": {
		"value": "33",
	},
	"12": {
		"value": "32",
	},
	"11": {
		"value": "30",
	},
	"10": {
		"value": "29",
	},
	"9": {
		"value": "28",
	},
	"8": {
		"value": "26",
		"1": map[string]string{
			"value": "27",
		},
	},
	"7": {
		"value": "24",
		"1": map[string]string{
			"value": "25",
		},
	},
	"6": {
		"value": "23",
	},
	"5": {
		"value": "21",
		"1": map[string]string{
			"value": "22",
		},
	},
	"4": {
		"value": "14",
		"1": map[string]string{
			"value": "16",
		},
		"2": map[string]string{
			"value": "17",
		},
		"3": map[string]string{
			"value": "18",
		},
		"4": map[string]string{
			"value": "19",
		},
	},
}
var (
	UserAgentFormat = "Mozilla/5.0 (Linux; Android {ANDROID_BUILD}; {ANDROID_MODEL} Build/{ANDROID_BUILD_M}; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.131 Mobile Safari/537.36 os/android model/{ANDROID_MODEL} build/{VERSION_CODE} osVer/{ANDROID_BUILD} sdkInt/{SDK_INT} network/2 BiliApp/{VERSION_CODE} mobi_app/android channel/{CHANNEL} Buvid/{BUVID} sessionID/{SESSION_ID} innerVer/{VERSION_CODE} c_locale/zh_CN s_locale/zh_CN disable_rcmd/0 {VERSION_NAME} os/android model/{ANDROID_MODEL} mobi_app/android build/{VERSION_CODE} channel/{CHANNEL} innerVer/{VERSION_CODE} osVer/{ANDROID_BUILD} network/2"
	BuildM          = "NRD90M.G955NKSU1AQDC"
	Channel         = "yingyongbao"
)
