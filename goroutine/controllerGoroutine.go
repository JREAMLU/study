package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var ch chan int
var wg *sync.WaitGroup

func main() {
	ch = make(chan int, 2)
	wg = &sync.WaitGroup{}
	fmt.Println("goroutine: ", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		ch <- 1
		wg.Add(1)
		go work(i)
	}
	wg.Wait()
}

func work(i int) {
	defer wg.Done()
	num := rand.Int31n(5)
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println("goroutine: ", runtime.NumGoroutine())
	fmt.Println("work: ", i)
	<-ch
}
