package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "http://mattock.plu.cn/test?name=kk&age=14&word=aa#ff"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("host: ", u.Host)
	fmt.Println("path: ", u.Path)
	fmt.Println("RawPath: ", u.RawPath)
	fmt.Println("rawQuery: ", u.RawQuery)
}
