package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 2)
	// var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(5 * time.Second)
			fmt.Println(i)
			ch <- i
		}(i)
	}

	for {
		select {
		case c := <-ch:
			fmt.Println("++++++++++++: ", c)
		default:
		}
	}
}
