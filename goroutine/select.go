package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 1)
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c <- 1
	c1 <- 2
	i := 3
	timeout := time.After(5 * time.Second)

	for {
		select {
		case s := <-c: //receive
			fmt.Println(s)
		case t := <-c1: //receive
			fmt.Println(t)
		case c2 <- i: //sent
			fmt.Println("sent")
		case <-timeout: //receive
			fmt.Println("You talk too much")
			return
		}
	}
}
