package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
	runtime.GC()
	var n runtime.MemStats
	runtime.ReadMemStats(&n)
	fmt.Printf("%d Kb\n", n.Alloc/1024)
}
