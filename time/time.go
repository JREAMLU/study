package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Location:", t.Location(), ":Time:", t)
	utc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println("Location:", utc, ":Time:", t.In(utc))

	zonename, offset := time.Now().In(utc).Zone()
	fmt.Println(zonename, offset)
}
