package main

import (
	"fmt"

	"github.com/mohong122/ip2region/binding/golang"
)

func main() {
	fmt.Println("err")
	region, err := ip2region.New("ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	sip := "119.75.218.70,127.0.0.1"
	ip, err := region.MemorySearch(sip)
	fmt.Println(ip, err)
	ip, err = region.BinarySearch(sip)
	fmt.Println(ip, err)
	ip, err = region.BtreeSearch(sip)
	fmt.Println(ip, err)
}
