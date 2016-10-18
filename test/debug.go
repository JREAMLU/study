package main

import "fmt"

func main() {
	a := 1
	b := 2
	tot := add(a, b)
	if tot > 3 {
		fmt.Println(tot)
	}
}

func add(a int, b int) int {
	return a + b
}
