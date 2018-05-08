package main

import (
	"fmt"
)

//go:generate msgp

// Foo foo
type Foo struct {
	Bar string  `msg:"bar"`
	Baz float64 `msg:"baz"`
}

func main() {
	fmt.Println("Nothing to see here yet!")
	d := marshal()
	f := unmarshal(d)
	fmt.Println("++++++++++++: ", f.Bar)
	fmt.Println("++++++++++++: ", f.Baz)
}

func marshal() []byte {
	foo1 := Foo{
		Bar: "bar",
		Baz: 1.2,
	}
	fmt.Printf("foo1: %v\n", foo1)

	// Here, we append two messages
	// to the same slice.
	data, _ := foo1.MarshalMsg(nil)
	fmt.Println("++++++++++++: ", string(data))

	return data
}

func unmarshal(d []byte) Foo {
	// Now we'll just decode them
	// in reverse:
	var foo1 Foo
	data, _ := foo1.UnmarshalMsg(d)
	fmt.Println("++++++++++++: ", string(data))

	// at this point, len(data) should be 0
	fmt.Println("len(data) =", len(data))

	fmt.Printf("foo1: %v\n", foo1)
	return foo1
}
