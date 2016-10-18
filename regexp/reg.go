package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `{"c":"jream","cd1":"jream"}`
	reg, err := regexp.Compile(`cd[0-9]*`)
	a := reg.FindString(str)
	fmt.Println(a, err)
}
