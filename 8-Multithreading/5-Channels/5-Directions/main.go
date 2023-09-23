package main

import "fmt"

func main() {
	channel := make(chan string)

	go recieve("Hello", channel)
	read(channel)
}

func recieve(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}
