package main

import "fmt"

func main() {
	ch1 := make(chan string)

	ch1 <- "hello"
	ch1 <- "world"

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}
