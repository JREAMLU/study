package main

import "fmt"

var all chan int

func main() {
	all = make(chan int, 3)
	all <- 1
	all <- 2
	all <- 3

	s := t1()
	for i := range s {
		// 推出标记
		if i == 3 {
			close(all)
		}
		fmt.Println("++++++++++++: ", i)
	}
}

func t1() <-chan int {
	return all
}
