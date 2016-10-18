package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 50000; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	t1 := time.Now()

	// go say("world") //开一个新的Goroutines执行
	say("hello") //当前Goroutines执行

	fmt.Println(time.Now().Sub(t1))
}
