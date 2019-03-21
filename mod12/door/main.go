package main

import (
	"fmt"

	"github.com/JREAMLU/j-kit/uuid"
	hello "github.com/JREAMLU/study/mod12"
)

func main() {
	fmt.Println("++++++++++++: ", hello.Hello())
	uuid.Generate()
}
