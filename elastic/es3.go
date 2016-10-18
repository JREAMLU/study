package main

import "study/core"

func main() {
	err := core.InitElastic("http://127.0.0.1:9200")
	if err != nil {
		panic(err)
	}
}
