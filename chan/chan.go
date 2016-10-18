package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(c chan int) {
	for i := 1; i <= 15; i++ {
		c <- i
		fmt.Println("发送一个:", i)
	}
}

func comsumer(c chan int, o chan bool) {
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println("handle接受一个: ", v)
				do(v)
			case <-time.After(7 * time.Second):
				fmt.Println("timeout")
				o <- true
				// break
			}
		}
	}()
	producer(c)
	<-o
}

func do(i int) {
	num := rand.Int63n(5)
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println("睡眠时间: ", i, "-", num)
}

func main() {
	c := make(chan int, 3)
	o := make(chan bool)
	comsumer(c, o)
}
