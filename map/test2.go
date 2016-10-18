package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["abc"] = "abc"
	fmt.Println(a["abc"])
}
