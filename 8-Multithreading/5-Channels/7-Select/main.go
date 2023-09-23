package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var counter int64 = 0

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			c1 <- Message{counter, "Hello from RabbitMQ"}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			c2 <- Message{counter, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from Kafka, ID: %d - %s\n", msg.id, msg.msg)
		case msg := <-c2:
			fmt.Printf("Redceived from RabbitMQ, ID: %d - %s\n", msg.id, msg.msg)
		case <-time.After(time.Second * 4):
			fmt.Println("Timeout")
		}
	}
}
