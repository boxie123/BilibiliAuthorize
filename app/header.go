package app

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"math"
	"strconv"
	"strings"
	"unicode"
)

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
