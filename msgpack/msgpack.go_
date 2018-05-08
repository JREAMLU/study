package main

import (
	"fmt"

	"github.com/vmihailenco/msgpack"
)

// Item item
type Item struct {
	Foo string
}

func main() {
	var item Item
	item.Foo = "bar"
	b, err := msgpack.Marshal(&item)
	fmt.Println("++++++++++++: ", b, err)
	var item2 Item
	err = msgpack.Unmarshal(b, &item2)
	fmt.Println("++++++++++++: ", item2, err)
}
