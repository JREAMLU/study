package main

import "fmt"

func main() {
	var s strs
	s.name = "1a"
	var s2 strs
	s2.name = "2a"
	tt(s, s2)
}

type strs struct {
	name string
}

func tt(str ...strs) {
	fmt.Println("++++++++++++: ", str)
}
