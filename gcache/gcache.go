package main

import (
	"fmt"

	"github.com/bluele/gcache"
)

func main() {
	cache := gcache.New(2).LFU().Build()
	cache.Set("Key", 1)
	cache.Set("Key2", 2)
	// val, err := cache.Get("Key")
	// if err != nil {
	// 	panic(err)
	// }
	vals := cache.GetALL()
	fmt.Println("++++++++++++: ", vals)
}
