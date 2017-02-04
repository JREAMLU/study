package main

import (
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var taskChannel chan int
var taskChannel = make(chan int, 100)
var total = 10

func main() {
	// serverV1()
	serverV2()
}

func serverV1() {
	s := time.Now().UnixNano()
	for i := 0; i < total; i++ {
		wg.Add(1)
		go processV1(i)
	}
	wg.Wait()
	e := time.Now().UnixNano()
	tt := e - s
	log.Printf("%d us", tt)
}

func processV1(task int) int {
	d := task + 1
	wg.Done()
	return d
}

func serverV2() {
	s := time.Now().UnixNano()
	consumerServerV2(1)

	go func() {
		for i := 1; i < total; i++ {
			taskChannel <- i
		}
	}()

	e := time.Now().UnixNano()
	tt := e - s
	log.Printf("%d us", tt)
}

func processV2(task int) int {
	d := task + 1
	log.Println(task)
	return d
}

func consumerServerV2(workNumber int) {
	for i := 0; i < workNumber; i++ {
		go func() {
			for {
				select {
				case task := <-taskChannel:
					processV2(task)
				case <-time.After(1 * time.Second):
					log.Println("超时")
				default:
					log.Println("阻塞")
				}
			}
		}()
	}
}
