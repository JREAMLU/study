package main

import (
	"fmt"
	"time"
)

func main() {
	ts := time.Now().Unix()
	ts2 := time.Now().UnixNano()
	ts3 := time.Now().UTC()
	fmt.Println(ts, ts2, ts3)
}
