package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"git.corp.plu.cn/mining/orion-server/core/com"
)

func main() {
	// str := "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0"
	// str := "Mozilla/5.0 (iPod; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5"
	// str := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36"
	// str := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	// str := "Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, like Gecko) Version/6.0.0.337 Mobile Safari/534.1+"
	// str := "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET4.0C; .NET4.0E)"
	str := "Lenovo-A858t_TD/S100 Linux/3.10.48 Android/4.4.4 Release/03.26.2013 Browser/AppleWebKit534.30 Mobile Safari/534.30 MBBMS/2.2"

	ua := com.ParseUserAgent(str)
	fmt.Println("ua: ", ua)
	spew.Dump("ua: ", ua)
}
