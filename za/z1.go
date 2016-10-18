package main

import "fmt"

type Box struct {
	width  string
	height string
}

type BoxList []Box

func main() {
	var b Box
	b.width = "10px"
	b.height = "20px"
	fmt.Println(b)

	bl := BoxList{
		Box{
			width:  "10px",
			height: "20px",
		},
		Box{
			width:  "30px",
			height: "40px",
		},
	}
	fmt.Println(bl)

	var bl2 BoxList
	bl2 = BoxList{
		Box{
			width:  "10px",
			height: "20px",
		},
		Box{
			width:  "30px",
			height: "40px",
		},
	}
	fmt.Println(bl2)
}
