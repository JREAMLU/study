package main

import "fmt"

func main() {
	var ch = make(chan int, 100)
	delta := 0

	for i := 0; i < delta; i++ {
		fmt.Println("<<<<")
		ch <- 1
	}

	for i := 0; i > delta; i-- {
		fmt.Println(">>>>")
		fmt.Println(<-ch)
	}
}
