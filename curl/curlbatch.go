package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/freshcn/async"
)

type aRequests struct {
	Method string
	UrlStr string
	Header map[string]string
	Raw    string
}

func aRollingCurl(r aRequests) (string, error) {
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

// 耗时操作1
func request1() interface{} {
	return Add(`{"name":"KII","age":24,"type":"a"}`)
}

// 耗时操作2
func request2() interface{} {
	return Add(`{"name":"KII","age":24,"type":"b"}`)
}

func Add(raw string) string {
	res, err := aRollingCurl(
		aRequests{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/test2.php",
			Header: map[string]string{
				"Content-Type": "application/json",
			},
			Raw: raw,
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func main() {
	// 建议程序开启多核支持
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 新建一个async对象
	async := async.New()

	// 添加request1异步请求,第一个参数为该异步请求的唯一logo,第二个参数为异步完成后的回调函数,回调参数类型为func()interface{}
	async.Add("request1", request1)
	// 添加request2异步请求
	async.Add("request2", request2)

	// 执行
	if chans, ok := async.Run(); ok {
		// 将数据从通道中取回,取回的值是一个map[string]interface{}类型,key为async.Add()时添加的logo,interface{}为该logo回调函数返回的结果
		res := <-chans
		// 这里最好判断下是否所有的异步请求都已经执行成功
		if len(res) == 2 {
			for k, v := range res {
				//do something
				fmt.Println(k, v)
			}
		} else {
			log.Println("async not execution all task")
		}
	}

	// 清除掉本次操作的所以数据,方便后续继续使用async对象
	async.Clean()
}
