package main

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	dispatch := kingpin.Flag("dispatch", "分发模式").Default("true").String()
	room := kingpin.Flag("room", "消费房间信息").Default("true").String()
	gift := kingpin.Flag("gift", "消费礼物").Default("true").String()

	kingpin.Parse()
	fmt.Println("++++++++++++: ", *dispatch)
	fmt.Println("++++++++++++: ", *room)
	fmt.Println("++++++++++++: ", *gift)

	b, err := strconv.ParseBool(*dispatch)
	spew.Dump(b)
	fmt.Println("++++++++++++: ", err)
}
