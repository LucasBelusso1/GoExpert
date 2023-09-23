package main

import "fmt"

func main() {
	channel := make(chan string) // Empty channel

	go func() {
		channel <- "Hello World!" // Full channel
	}()

	msg := <-channel // Channel gets empty again
	fmt.Println(msg)
}
