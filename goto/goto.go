package main

import "fmt"

func main() {
	i := 0
A:
	i++
	if i < 3 {
		goto A
	}
	fmt.Println(i)
}
