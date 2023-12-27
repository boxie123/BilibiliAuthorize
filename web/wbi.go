package web

import (
	"encoding/json"
	"github.com/boxie123/BilibiliAuthorize/utils"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ApiNav = "https://api.bilibili.com/x/web-interface/nav" // 导航栏用户信息, 用于获取 img_url 和 sub_url
)

var (
	mixinKeyEncTab = [...]int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
	cache          sync.Map
	lastUpdateTime time.Time
)

// 从伪装的图片链接中提取 token
func getTokenFromUrl(wbiUrl string) string {
	urlSli := strings.Split(wbiUrl, "/")
	imgName := urlSli[len(urlSli)-1]
	key := strings.Split(imgName, ".")[0]
	return key
}

// UpdateWbiKey 更新缓存中的 img_key 和 sub_key
func UpdateWbiKey() {
	resp, err := http.Get("https://api.bilibili.com/x/web-interface/nav")
	if err != nil {
		log.Printf("Error: %s", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	var result NavResp
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	imgUrl := result.Data.WbiImg.ImgURL
	subUrl := result.Data.WbiImg.SubURL
	cache.Store("imgKey", getTokenFromUrl(imgUrl))
	cache.Store("subKey", getTokenFromUrl(subUrl))
	lastUpdateTime = time.Now()
}

// 从缓存中读取 WbiKey
func getWbiKeysCached() (string, string) {
	imgKeyI, _ := cache.Load("imgKey")
	subKeyI, _ := cache.Load("subKey")
	if imgKeyI == "" || subKeyI == "" {
		UpdateWbiKey()
	}
	imgKeyI, _ = cache.Load("imgKey")
	subKeyI, _ = cache.Load("subKey")
	return imgKeyI.(string), subKeyI.(string)
}

// 按映射表生成 MixinKey
func getMixinKey(orig string) string {
	var str strings.Builder
	for _, v := range mixinKeyEncTab {
		if v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	return str.String()[:32]
}

// ParamSign 为参数添加 wts 和 w_rid
func ParamSign(param map[string]interface{}) map[string]interface{} {
	_, ok := param["wts"]
	if !ok {
		param["wts"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	paramQuery := utils.ParamEncode(param)
	// 每10分钟更新一次 WbiKey
	if time.Since(lastUpdateTime).Minutes() > 10 {
		UpdateWbiKey()
	}
	imgKey, subKey := getWbiKeysCached()
	wRid := utils.Md5(paramQuery + getMixinKey(imgKey+subKey))
	param["w_rid"] = wRid
	return param
}
