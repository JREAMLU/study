package main

import (
	l4g "git.corp.plu.cn/plugo/log4go"
)

func main() {
	l4g.LoadConfiguration("./log.xml")
	defer l4g.Close()
	l4g.Warn("abc")
}
