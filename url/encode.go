package main

import (
	"fmt"
	"net/url"
)

func main() {
	//urlencode
	param := "name=发&age=第三方"
	encode := url.QueryEscape(param)
	fmt.Println(encode)
	//urldecode
	unencode, err := url.QueryUnescape(encode)
	fmt.Println(unencode, err)
}
