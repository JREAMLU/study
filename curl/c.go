package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func httpGet() {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}
	requsetUrl := "http://localhost/study/curl/get.php?a=1&b=2"
	resp, err := c.Get(requsetUrl)
	if err != nil {
		fmt.Println("request err:", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err:", err)
	}

	fmt.Println(string(body))
}

func main() {
	httpGet()
}
