package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// var data = `
// word:
//     words:
//       - <content> - abc
//       - b
// `

var data = `
word:
    words:
      - <content> - 您订阅的<name>开播了!订阅获取更多主播开播提醒
      - <content> - <name>开播了!过这村就没这店, 下次还不知道要等多久啊~
      - <content>在TA的直播间提起了你, 快来参加<name>的讨论
`

type AT struct {
	Word struct {
		Words []string `yaml:"words"`
	} `yaml:"word"`
}

func aread(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func main() {
	var t AT

	err := yaml.Unmarshal([]byte(data), &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", t.Word.Words)

}
