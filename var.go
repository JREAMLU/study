package main

import "fmt"

var list []struct {
	Name    string
	Address string
}

type man []struct {
	Name    string
	Address string
}

func main() {
	// var numbers map[string]int
	// numbers := make(map[string]interface{})
	// numbers["one"] = 1  //赋值
	// numbers["ten"] = 10 //赋值
	// numbers["three"] = []string{"abc"}
	// fmt.Println(numbers)
	// list.Name = "jream"
	// list.Address = "ppxc"
	// fmt.Println(list)

	// m := man{{Name: "jream", Address: "shanghai"}, {Name: "jream", Address: "shanghai"}}

	list = []struct {
		Name    string
		Address string
	}{
		{Name: "jream", Address: "shanghai"},
		{Name: "jream", Address: "shanghai"},
	}

	fmt.Println(list)
}
