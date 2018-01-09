package main

import (
	"fmt"

	_ "github.com/JREAMLU/study/interface/plugins/p1"
	// _ "github.com/JREAMLU/study/interface/plugins/p2"
	"github.com/JREAMLU/study/interface/server"
)

func main() {
	str, _ := server.MG.Start("gogogo")
	fmt.Println("++++++++++++: ", str)
}
