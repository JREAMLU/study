package main

import "fmt"

type Gointo interface {
	Read(str string)
	Write(id int)
}

type Men struct {
	id   int
	name string
}

type Women struct {
	id   int
	name string
}

func (m *Men) Read(str string) {
	fmt.Println("Men Read: ", str)
}

func (m *Men) Write(id int) {
	fmt.Println("Men Write: ", id)
}

func (w *Women) Read(str string) {
	fmt.Println("Women Read: ", str)
}

func (w *Women) Write(id int) {
	fmt.Println("Women Write: ", id)
}

func main() {
	var g Gointo
	g = new(Men)
	g = new(Women)
	str := "abc"
	g.Read(str)
}
