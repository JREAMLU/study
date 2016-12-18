package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			buf := make([]byte, 1<<20)
			runtime.Stack(buf, false)

			spew.Dump(x)
			// Logger(x)
		}
	}()
	var s []string
	s = append(s, "a")
	fmt.Println(s[10])
}

func Logger(v interface{}) {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	_, err = conn.Write([]byte(v.(string)))
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		log.Printf("%v Fatal error %v", os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(string(buf[0:n]))
	os.Exit(0)

}

func Critical(v ...interface{}) {
	logs.Critical(generateFmtStr(len(v)), v...)
}
func Li function() {
    if LevelCritical > bl.level {
        return
    }
    bl.writeMsg(LevelCritical, format, v...)
}
func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
