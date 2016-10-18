package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	a := "abc"
	b := crc32.ChecksumIEEE([]byte(a))
	fmt.Println(b)
}
