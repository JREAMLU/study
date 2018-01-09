package main

import (
	"fmt"

	// _ "github.com/JREAMLU/study/interface/plugins/p1"
	_ "github.com/JREAMLU/study/interface/plugins/p2"
	"github.com/JREAMLU/study/interface/server"
)

func main() {
	// 根据参数导入包

	// cmd := "default"
	// cmd := "p1"
	cmd := "p2"
	str, _ := server.MG[cmd].Start("gogogo")
	fmt.Println("++++++++++++: ", str)
}
