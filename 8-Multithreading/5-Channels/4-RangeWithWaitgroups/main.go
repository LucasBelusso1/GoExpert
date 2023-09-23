package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go publish(channel, &waitGroup)
	go reader(channel)
	waitGroup.Wait()
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
