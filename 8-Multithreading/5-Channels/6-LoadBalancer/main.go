package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// Initialize workers
	for i := 0; i < 1000; i++ {
		go worker(i, channel)
	}

	// Request simulation
	for i := 0; i < 1000000; i++ {
		channel <- i
	}
}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
