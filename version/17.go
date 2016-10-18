package main

import "github.com/davecgh/go-spew/spew"

func main() {
	var str interface{}
	spew.Dump(str)
	spew.Dump(str.(string))
}
