package utils

import (
	"testing"
)

func TestParamEncode(t *testing.T) {
	param := map[string]string{
		"foo": "one one four",
		"bar": "五一四",
		"baz": "1919810",
	}
	ps := ParamEncode(param)
	if ps != "bar=%E4%BA%94%E4%B8%80%E5%9B%9B&baz=1919810&foo=one%20one%20four" {
		t.Errorf("参数编码错误: ps = %s", ps)
	}
}
