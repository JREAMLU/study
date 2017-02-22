package main

import (
	"fmt"
	"time"
)

func main() {
	t := 1486548310
	t = 1486549191
	// m := 1486544752464946010
	tm := time.Unix(int64(t), 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05 PM"))

	fmt.Println(time.Now().UnixNano())
}
