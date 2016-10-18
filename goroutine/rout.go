package main

import (
	"fmt"
	"time"
)

/*
func a() {
	fmt.Println("a")
	w := make(chan int)
	b(w)
	c()
	<-w
}
*/

func b() {
}

func main() {
	c := make(chan int)
	d := make(chan int)
	e := make(chan int)
	f := make(chan int)

	go func(c chan int) {
		b()
		time.Sleep(3 * time.Second)
		c <- 1
	}(c)

	go func(d chan int) {
		b()
		time.Sleep(3 * time.Second)
		d <- 2
	}(d)

	go func(e chan int) {
		b()
		time.Sleep(3 * time.Second)
		e <- 3
	}(e)

	go func(f chan int) {
		b()
		time.Sleep(3 * time.Second)
		f <- 4
	}(f)

	o := make(chan bool)
	go func() {
		i := 0
		for {
			if i >= 4 {
				o <- true
			}
			select {
			case c := <-c:
				fmt.Println(c)
				i++
			case d := <-d:
				fmt.Println(d)
				i++
			case e := <-e:
				fmt.Println(e)
				i++
			case f := <-f:
				fmt.Println(f)
				i++
			default:
			}
		}
	}()
	<-o

	fmt.Println("done")
	// fmt.Println(<-c)
	// fmt.Println(<-d)
}
