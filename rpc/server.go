package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type ServerCommend int

// type Result struct {
// 	out string
// }

func (s *ServerCommend) Doshell(cmd string, reply *string) error {
	fmt.Println(cmd)
	id := make(chan int)
	go baseCommand(cmd, id)
	pid := <-id
	*reply = strconv.Itoa(pid)
	return nil
}

func main() {
	sc := new(ServerCommend)
	rpc.Register(sc)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func baseCommand(shell string, id chan int) (int, error) {
	cmd := exec.Command("cmd", "/C", shell)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return cmd.Process.Pid, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return cmd.Process.Pid, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return cmd.Process.Pid, err
	}

	pid := cmd.Process.Pid
	id <- pid

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		return cmd.Process.Pid, err
	}

	if len(bytesErr) != 0 {
	}

	bytes, err := ioutil.ReadAll(stdout)
	fmt.Println(string(bytes))

	if err != nil {
		return cmd.Process.Pid, err
	}

	if err := cmd.Wait(); err != nil {
		rex, _ := regexp.Compile("terminated")
		ok := rex.Match([]byte(err.Error()))
		if !ok {
		}

		return pid, err
	}

	return pid, nil
}
