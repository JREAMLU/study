package main

import (
	"fmt"

	"github.com/xluohome/phonedata"
)

func main() {
	pr, err := phonedata.Find("1580000000")
	if err != nil {
		panic(err)
	}
	fmt.Print(pr)
}

