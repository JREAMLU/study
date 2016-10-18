package main2

import (
	"fmt"
	"os/exec"
)

func run() {
	cmd := exec.Command("cmd", "/C", "ping 127.0.0.1 > test.txt")
	err := cmd.Start()
	if err != nil {
		fmt.Println("exec failed")
	}
	/*
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("StdoutPipe: " + err.Error())
			return
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println("StderrPipe: ", err.Error())
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Start: ", err.Error())
			return
		}

		bytesErr, err := ioutil.ReadAll(stderr)
		if err != nil {
			fmt.Println("ReadAll stderr: ", err.Error())
			return
		}

		if len(bytesErr) != 0 {
			fmt.Printf("stderr is not nil: %s", bytesErr)
			return
		}

		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll stdout: ", err.Error())
			return
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("Wait: ", err.Error())
			return
		}

		fmt.Printf("stdout: %s", bytes)
	*/
}

func main() {
	run()
}
