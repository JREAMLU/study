package main

import (
	"fmt"

	"github.com/deepzz0/go-com/useragent"
	// "github.com/JREAMLU/core/useragent"
)

func main() {
	var str = `Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36 QIHU 360EE`
	// var str = `Mozilla/5.0 (iPhone 5ATT; CPU iPhone OS 9_3_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/6.0 MQQBrowser/6.8.1 Mobile/13F69 Safari/8536.25 MttCustomUA/2`
	// var str = `Mozilla/5.0 (Linux; Android 5.1; 1501_M02 Build/LMY47D) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.94 Mobile Safari/537.36 360 Aphone Browser (100.2.0)`
	// var str = `Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36`

	agent := useragent.ParseByString(str)
	fmt.Println(agent)
}
