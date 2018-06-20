package main

import "fmt"

func main() {
	t1("luj", func(last string) {
		fmt.Println("++++++++++++: ", last)
	})
}

func t1(name string, add func(last string)) {
	if name == "luj" {
		add(name)
	}
}
