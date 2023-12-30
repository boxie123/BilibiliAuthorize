# BilibiliAuthorize
 Go实现的Bilibili Api签名

## 使用
```cmd
go get -u github.com/boxie123/BilibiliAuthorize/web
```

```cmd
go get -u github.com/boxie123/BilibiliAuthorize/app
```
按需引入

## 示例
```go
package main

import (
	"fmt"
	"github.com/boxie123/BilibiliAuthorize/web"
	"github.com/boxie123/BilibiliAuthorize/app"
)

func main() {
	param := map[string]string{
		"foo": "114",
		"bar": "514",
		"zab": "1919810",
	}
	param1 := web.ParamSign(param)
	param2 := app.ParamSign(param)
	fmt.Println(param1, param2)
}
```

## WARNING

- 本项目遵守 GPLv3 协议，如需使用、转载必须保留版权和许可声明并公开源码
- 请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！请勿滥用，本项目仅用于学习和测试！
- 利用本项目提供的接口、文档等造成不良影响及后果与本人无关
- 本项目为开源项目，不接受任何形式的催单和索取行为，更不容许存在付费内容
