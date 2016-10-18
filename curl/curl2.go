package main

import (
	"encoding/json"
	"fmt"

	"github.com/JREAMLU/core/curl"
)

func main() {
	url := "http://172.16.9.221:8080/collect?v=1&_v=mithril&cid=1230619451.1469690450&tid=1321431-1&t=pageview&dl=http%3A%2F%2Fmattock.plu.cn%2Ftest%2F&ul=zh-cn&de=0&dt=test&sd=24-bit&sr=1920x1080&vp=147x677&fl=22.0%20r0&_s=1"
	res, err := getHttp(url)
	fmt.Println("res: ", res)
	fmt.Println("err: ", err)
}

func getHttp(url string) (map[string]interface{}, error) {
	res, err := curl.RollingCurl(
		curl.Requests{
			Method: "GET",
			UrlStr: url,
		},
	)
	if err != nil {
		return nil, err
	}
	var result = make(map[string]interface{})
	json.Unmarshal([]byte(res), &result)

	return result, nil
}
