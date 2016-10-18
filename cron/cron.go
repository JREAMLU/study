package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	// "github.com/robfig/cron"
)

func main() {
	kafkaConnNum := `netstat -an | grep 9092 | wc -l`
	hmConnNum := `netstat -an | grep 8080 | wc -l`
	kafkaConn := `netstat -n | grep 9092 | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'`
	hmConn := `netstat -an | grep 8090 | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'`

	kafka, _ := connects(kafkaConn)
	kafka_num, _ := connNum(kafkaConnNum)
	collect, _ := connects(hmConn)
	collect_num, _ := connNum(hmConnNum)

	health := make(map[string]interface{})
	health["kafka_num"] = kafka_num
	health["kafka"] = kafka
	health["collect_num"] = collect_num
	health["collect"] = collect

	fmt.Println(health)

	// c := cron.New()
	// c.AddFunc("*/10 * * * * *", func() {
	// 	fmt.Println("每一分钟执行:", time.Now().Format("2006-01-02 15:04:05"))
	// 	//ps
	// })
	// c.Start()
	// fmt.Println("检查: ", c.Entries())
	//
	// for {
	// }
}

//connects 连接数
func connects(shell string) (map[string]string, error) {
	str, err := LocalCommand(shell)
	if err != nil {
		return nil, err
	}

	list := make(map[string]string)

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		word := strings.Split(line, " ")
		if len(word) > 1 {
			list[word[0]] = word[1]
		}

	}

	return list, nil
}

//connTotal 连接总数
func connNum(shell string) (string, error) {
	str, err := LocalCommand(shell)
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, " ", "", -1)
	num := strings.Replace(str, "\n", "", -1)

	return num, nil
}

//LocalCommand 本机命令
func LocalCommand(shell string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", shell)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return "", err
	}

	if len(bytesErr) != 0 {
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return string(bytes), nil
}
