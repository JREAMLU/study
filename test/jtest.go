package jtest

import "fmt"

func Add(i int) int {
	fmt.Println("add")
	if i == 1 {
		fmt.Println(i)
		return 1
	} else {
		fmt.Println(i)
		return 2
	}
}

func Update(i int) int {
	fmt.Println("update")
	if i == 1 {
		fmt.Println(i)
		return 1
	} else {
		fmt.Println(i)
		return 2
	}
}
