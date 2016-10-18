package main

import "fmt"

func main() {
	var a = make(map[string]interface{})
	var b = make(map[string]interface{})
	var d = make(map[string]interface{})

	a["name"] = "abc"
	b["addr"] = "china"
	b["name"] = "def"

	c := MapMerge(a, b, d)
	fmt.Println(c)
}

func MapMerge(ms ...map[string]interface{}) map[string]interface{} {
	var nm = make(map[string]interface{})
	for _, m := range ms {
		for k, v := range m {
			nm[k] = v
		}
	}
	return nm
}
