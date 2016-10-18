package main

import (
	"fmt"

	log "github.com/thinkboy/log4go"
)

func main() {
	// log.LoadConfiguration("example.xml")
	filename := "abc"
	log.AddFilter("file", log.INFO, log.NewFileLogWriter(filename, false))

	flw := log.NewFileLogWriter(filename, false)
	flw.SetFormat("[%D %T] [%L] (%S) %M")
	flw.SetRotate(false)
	flw.SetRotateSize(0)
	flw.SetRotateLines(0)
	flw.SetRotateDaily(false)
	log.AddFilter("file", log.INFO, flw)
	fmt.Println(log.INFO)

	s := "abc"
	log.Info("记录日志 %s", s)

	// for i := 0; i < 100; i++ {
	// }

	for {

	}

}
