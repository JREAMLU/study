package main

import (
	"fmt"
	"strings"
)

func main() {
	// str := `<content> - 您订阅的<name>开播了!订阅获取更多主播开播提醒`
	str := `<content> - 您订阅的开播了!订阅获取更多主播开播提醒`
	content := "go"
	name := "jream"
	str = strings.Replace(str, "<content>", content, -1)
	str = strings.Replace(str, "<name>", name, -1)
	fmt.Println(str)
}
