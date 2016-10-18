package main1

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()
}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

func testTimer3() {
	fmt.Println("test timer3")
}

func timer1() {
	timer1 := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}

func timer2() {
	timer2 := time.NewTicker(2 * time.Second)
	timer3 := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-timer2.C:
			testTimer2()
		case <-timer3.C:
			testTimer3()
		}
	}
}

func cha(ch chan int) {
	fmt.Println("cha: ", ch)
	ch <- 1
}

func cha2(ch chan int) {
	fmt.Println("cha2: ", ch)
	ch <- 2
}

func main() {
	// go timer1()
	// timer2()
	/*
		ch := make(chan int)
		go cha(ch)
		go cha2(ch)
		x, y := <-ch, <-ch
		fmt.Println("x: ", x)
		fmt.Println("y: ", y)

		pack.ReadKafka()
	*/

	in := bytes.NewBuffer(nil)
	cmd := exec.Command("cmd", "/C")
	cmd.Stdin = in
	// go func() {
	// in.WriteString("echo hello world > test.txt\n")
	// in.WriteString("exit\n")
	// }()
	in.WriteString("ping 127.0.0.1")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
	buf, _ := cmd.Output()
	fmt.Fprintf(os.Stdout, "Result: %s", buf)
	// cmd := exec.Command("cmd")
	// buf, _ := cmd.Output()
	// fmt.Println(buf)
	// fmt.Fprintf(os.Stdout, "Result: %s", buf)
}
