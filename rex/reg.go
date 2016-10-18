package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		err := errors.New("signal: terminated")
		// str := "signal: terminated"
		str := err.Error()
		rex, _ := regexp.Compile("terminated")
		is := rex.Match([]byte(str))
		fmt.Println(is)
	*/

	// rex, _ := regexp.Compile("\\r|/\n")
	// dt := rex.ReplaceAllString(title, "")
	// fmt.Println(dt)
	title := "test\\r\\nabc\\rsdfa\\n"

	title = strings.Replace(title, "\\r", "", -1)
	title = strings.Replace(title, "\\n", "", -1)

	fmt.Println(title)

}
