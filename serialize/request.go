package main

import (
	"encoding/json"
	"fmt"
)

type AyiName struct {
	Data struct {
		Urls []struct {
			LongURL string `json:"long_url"`
			IP      string `json:"IP"`
		} `json:"urls"`
		Timestamp int64  `json:"timestamp"`
		Sign      string `json:"sign"`
	} `json:"data"`
}

func main() {
	// var an AyiName
	// var urls []map[string]string
	// urls[0]["long_url"] = "http://o9d.cn"
	// urls[0]["IP"] = "127.0.0.1"
	// urls[1]["long_url"] = "http://huiyimei.com"
	// urls[1]["IP"] = "192.168.1.1"
	// an.Data.Timestamp = time.Now().Unix()
	// an.Data.Urls = urls

	var requestParams = make(map[string]interface{})
	var data = make(map[string]interface{})
	var urls []map[string]string
	urls = append(urls, map[string]string{"long_url": "http://o9d.cn", "IP": "127.0.0.1"})
	urls = append(urls, map[string]string{"long_url": "http://huiyimei.net", "IP": "192.168.1.1"})
	data["urls"] = urls
	requestParams["data"] = data

	js, _ := json.Marshal(requestParams)

	fmt.Println("==================================", string(js))
}
