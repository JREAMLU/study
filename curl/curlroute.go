package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/profile"
)

type Requestsroute struct {
	Method string
	UrlStr string
	Header map[string]string
	Raw    string
}

func RollingCurlroute(r Requestsroute) (string, error) {
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

func Aaa(raw string) string {
	res, err := RollingCurlroute(
		Requestsroute{
			Method: "POST",
			UrlStr: "http://localhost/study/curl/test.php",
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

func rou() {
	ra := `{"name":"KII","age":24,"type":"a"}`
	rb := `{"name":"KII","age":24,"type":"b"}`
	rc := `{"name":"KII","age":24,"type":"c"}`
	rd := `{"name":"KII","age":24,"type":"d"}`
	re := `{"name":"KII","age":24,"type":"e"}`
	rf := `{"name":"KII","age":24,"type":"f"}`
	rg := `{"name":"KII","age":24,"type":"g"}`
	rh := `{"name":"KII","age":24,"type":"h"}`

	quita := make(chan string)
	quitb := make(chan string)
	quitc := make(chan string)
	quitd := make(chan string)
	quite := make(chan string)
	quitf := make(chan string)
	quitg := make(chan string)

	go func(quita chan string) {
		a := Aaa(ra)
		quita <- a
	}(quita)

	go func(quitb chan string) {
		b := Aaa(rb)
		quitb <- b
	}(quitb)

	go func(quitc chan string) {
		c := Aaa(rc)
		quitc <- c
	}(quitc)

	go func(quitd chan string) {
		d := Aaa(rd)
		quitd <- d
	}(quitd)

	go func(quite chan string) {
		e := Aaa(re)
		quite <- e
	}(quite)

	go func(quitf chan string) {
		f := Aaa(rf)
		quitf <- f
	}(quitf)

	go func(quitg chan string) {
		g := Aaa(rg)
		quitg <- g
	}(quitg)

	h := Aaa(rh)

	a, b, c, d, e, f, g := <-quita, <-quitb, <-quitc, <-quitd, <-quite, <-quitf, <-quitg
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
}

func xou() {
	ra := `{"name":"KII","age":24,"type":"a"}`
	rb := `{"name":"KII","age":24,"type":"b"}`
	rc := `{"name":"KII","age":24,"type":"c"}`
	rd := `{"name":"KII","age":24,"type":"d"}`
	re := `{"name":"KII","age":24,"type":"e"}`
	rf := `{"name":"KII","age":24,"type":"f"}`
	rg := `{"name":"KII","age":24,"type":"g"}`
	rh := `{"name":"KII","age":24,"type":"h"}`

	quit := make(chan string)

	go func(quit chan string) {
		a := Aaa(ra)
		quit <- a
	}(quit)

	go func(quit chan string) {
		b := Aaa(rb)
		quit <- b
	}(quit)

	go func(quit chan string) {
		c := Aaa(rc)
		quit <- c
	}(quit)

	go func(quit chan string) {
		d := Aaa(rd)
		quit <- d
	}(quit)

	go func(quit chan string) {
		e := Aaa(re)
		quit <- e
	}(quit)

	go func(quit chan string) {
		f := Aaa(rf)
		quit <- f
	}(quit)

	go func(quit chan string) {
		g := Aaa(rg)
		quit <- g
	}(quit)

	h := Aaa(rh)

	for {
		select {
		case r := <-quit:
			fmt.Println("==: ", r)
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	fmt.Println("h: ", h)
}

func cou() {
	ra := `{"name":"KII","age":24,"type":"a"}`
	rb := `{"name":"KII","age":24,"type":"b"}`
	rc := `{"name":"KII","age":24,"type":"c"}`
	rd := `{"name":"KII","age":24,"type":"d"}`
	re := `{"name":"KII","age":24,"type":"e"}`
	rf := `{"name":"KII","age":24,"type":"f"}`
	rg := `{"name":"KII","age":24,"type":"g"}`
	rh := `{"name":"KII","age":24,"type":"h"}`

	a := Aaa(ra)
	b := Aaa(rb)
	c := Aaa(rc)
	d := Aaa(rd)
	e := Aaa(re)
	f := Aaa(rf)
	g := Aaa(rg)
	h := Aaa(rh)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)

}

func ao() {
	t1 := time.Now()
	ra := `{"name":"KII","age":24,"type":"a"}`
	rb := `{"name":"KII","age":24,"type":"b"}`
	rc := `{"name":"KII","age":24,"type":"c"}`
	rd := `{"name":"KII","age":24,"type":"d"}`
	re := `{"name":"KII","age":24,"type":"e"}`
	rf := `{"name":"KII","age":24,"type":"f"}`
	rg := `{"name":"KII","age":24,"type":"g"}`
	rh := `{"name":"KII","age":24,"type":"h"}`

	go Aaa(ra)
	go Aaa(rb)
	go Aaa(rc)
	go Aaa(rd)
	go Aaa(re)
	go Aaa(rf)
	go Aaa(rg)
	Aaa(rh)

	fmt.Println(time.Now().Sub(t1))
	select {}
}

func co() {
	t1 := time.Now()
	ra := `{"name":"KII","age":24,"type":"a"}`
	rb := `{"name":"KII","age":24,"type":"b"}`
	rc := `{"name":"KII","age":24,"type":"c"}`
	rd := `{"name":"KII","age":24,"type":"d"}`
	re := `{"name":"KII","age":24,"type":"e"}`
	rf := `{"name":"KII","age":24,"type":"f"}`
	rg := `{"name":"KII","age":24,"type":"g"}`
	rh := `{"name":"KII","age":24,"type":"h"}`

	Aaa(ra)
	Aaa(rb)
	Aaa(rc)
	Aaa(rd)
	Aaa(re)
	Aaa(rf)
	Aaa(rg)
	Aaa(rh)

	fmt.Println(time.Now().Sub(t1))
	select {}
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	rou()
}
