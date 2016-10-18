package main

import "fmt"

func a(i int, j int) int {
	return i + j
}

func main() {
	i := 1
	j := 2
	rs := a(i, j)
	fmt.Println(rs)
}
