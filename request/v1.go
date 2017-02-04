package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go send()
	}
	select {}
}

func send() {
	beginTime := time.Now().UnixNano()
	res, err := RollingCurl(
		Requests{
			Method: "POST",
			UrlStr: "http://localhost/longzhu/study/server/go.php",
			Header: map[string]string{
				"Content-Type": "application/json;charset=UTF-8;",
			},
			Raw: string(""),
		},
	)

	fmt.Println("res: ", res)
	fmt.Println("err: ", err)

	endTime := time.Now().UnixNano()
	takeTime := endTime - beginTime
	log.Printf("time: %vms beginTime:%v endTime:%v ", takeTime/1000000, beginTime, endTime)
}
