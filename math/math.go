package main

import "fmt"

func main() {
	a := -1
	fmt.Println(IAbs(a))
}

func IAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IAbs32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
func IAbs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
