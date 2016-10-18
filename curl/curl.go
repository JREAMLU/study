package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpGet() {
	requsetUrl := "http://localhost/study/curl/get.php?a=1&b=2"
	resp, err := http.Get(requsetUrl)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPost() {
	requsetUrl := "http://localhost/study/curl/servera.php"
	raw := `{"name":"KII","age":24}`
	resp, err := http.Post(requsetUrl,
		"application/json",
		strings.NewReader(raw))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	requsetUrl := "http://localhost/study/curl/servera.php"
	var formdata = make(map[string][]string)
	formdata["name"] = []string{"KII"}
	formdata["age"] = []string{"24"}
	resp, err := http.PostForm(requsetUrl, formdata)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

type Requests struct {
	Method string
	UrlStr string
	Header map[string]string
	Raw    string
}

func RollingCurl(r Requests) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		r.Method,
		r.UrlStr,
		strings.NewReader(r.Raw),
	)

	if err != nil {
		return "", err
	}

	for hkey, hval := range r.Header {
		req.Header.Set(hkey, hval)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	res, err := RollingCurl(
		Requests{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/servera.php",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Raw: `{"name":"KII","age":24}`,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	var result = make(map[string]interface{})
	json.Unmarshal([]byte(res), &result)
	for k, v := range result {
		fmt.Printf("%s: %v \n", k, v)
	}
}
