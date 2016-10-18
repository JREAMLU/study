package main

import "fmt"

func main() {
	a := []int{1, 2}
	if len(a) > 1 {
		fmt.Println(a[1])
	}
}
