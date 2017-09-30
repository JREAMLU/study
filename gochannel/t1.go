package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	go func() {
		for {
			fmt.Println("Waiting for job...")
			select {
			// Receive from channel
			case j := <-Jobs:
				fmt.Println("worker", id, "started  job", j)
				time.Sleep(time.Second)
				fmt.Println("worker", id, "finished job", j)
				Results <- true
			}
		}
	}()
}

const channelLength = 3

var (
	Jobs    chan int
	Results chan bool
)

func main() {
	Jobs = make(chan int, channelLength)
	Results = make(chan bool, channelLength)

	// Start worker goroutines
	for i := 0; i < channelLength; i++ {
		worker(i)
	}

	// Send to channel
	time.Sleep(time.Second)
	for j := 0; j < channelLength; j++ {
		Jobs <- j
	}
	close(Jobs)

	for len(Jobs) != 0 || len(Results) != channelLength {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Complete main")
}
