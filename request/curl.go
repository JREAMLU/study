package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Requests struct {
	Method string
	UrlStr string
	Header map[string]string
	Raw    string
}

//RollingCurl http请求url
func RollingCurl(r Requests) (string, error) {
	client := &http.Client{
		Timeout: 3600 * time.Second,
	}

	req, err := http.NewRequest(
		r.Method,
		r.UrlStr,
		strings.NewReader(r.Raw),
	)
	fmt.Println(req, err)

	if err != nil {
		return "", err
	}

	for hkey, hval := range r.Header {
		req.Header.Set(hkey, hval)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
