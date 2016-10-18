package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	buffer.WriteString("abc")
	buffer.WriteString("def")

	fmt.Println(buffer.String())
}
