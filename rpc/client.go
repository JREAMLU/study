package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/robfig/cron"
)

func main() {
	CronRun()
}

func CronRun() {
	schedule, _ := cron.Parse("*/10 * * * * *")
	entry := &cron.Entry{
		Schedule: schedule,
	}

	var c cron.Cron
	c.Pentries = append(c.Pentries, entry)

	now := time.Now().Local()

	for _, entry := range c.Pentries {
		entry.Next = entry.Schedule.Next(now)
	}

	fmt.Println("CronRun c.Pentries[0].Next: %d  pid: %d", c.Pentries[0].Next, os.Getpid())

	for {
		var effective time.Time
		effective = c.Pentries[0].Next
		fmt.Println("下次执行时间: %v", effective)
		select {
		case now = <-time.After(effective.Sub(now)):
			for _, e := range c.Pentries {
				if e.Next != effective {
					break
				}
				fmt.Println("CronRun effective.Sub(now): %v", effective)

				DoRPC()

				e.Prev = e.Next
				e.Next = e.Schedule.Next(effective)
			}
			continue
		}
	}
}

func DoRPC() {
	client, err := rpc.Dial("tcp", "172.16.9.230:1234")
	defer client.Close()
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("ServerCommend.Doshell", `ping 127.0.0.1`, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("ServerCommend: %s \n", reply)
}

//LocalCommand 本机命令
func LocalCommand(shell string) (int, error) {
	cmd := exec.Command("/bin/sh", "-c", shell)
	fmt.Println("shell: %v", shell)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: %v", err.Error())
		return cmd.Process.Pid, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: %v", err.Error())
		return cmd.Process.Pid, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: %v", err.Error())
		return cmd.Process.Pid, err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: %v", err.Error())
		return cmd.Process.Pid, err
	}

	if len(bytesErr) != 0 {
		fmt.Println("stderr is not nil: %s", bytesErr)
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: %v", err.Error())
		return cmd.Process.Pid, err
	}

	pid := cmd.Process.Pid
	fmt.Println("stdout: %s", bytes)

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: %v", err.Error())

		rex, _ := regexp.Compile("terminated")
		ok := rex.Match([]byte(err.Error()))
		if !ok {
			alarmMsg := "退出信号:<" + err.Error() + "> 进程pid:<" + strconv.Itoa(pid) + "> "
			fmt.Println("警报短信内容: %s ", alarmMsg)
		}

		return pid, err
	}

	return pid, nil
}

/*
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	cmd := "ping 127.0.0.1 -t"
	var reply string
	err = client.Call("ServerCommend.Doshell", cmd, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("ServerCommend: %s : %s \n", cmd, reply)

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	service := os.Args[1]

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	cmd := "ping 127.0.0.1 -t"
	var reply string
	err = client.Call("RpcServer.Doshell", `{"Id":1,"Shell":"ping 127.0.0.1","Name":"aa"}`, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("ServerCommend: %s : %s \n", cmd, reply)

}
*/
