package main

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	a := os.Environ()
	spew.Dump(a)
}
