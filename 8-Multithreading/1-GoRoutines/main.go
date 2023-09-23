package main

import (
	"fmt"
	"time"
)

func main() {
	// Normal execution, A stops before B starts:
	// task("A")
	// task("B")

	// Paralel (but doesn't work):
	go task("A")
	go task("B")

	// Paralel (working):
	go task("A")
	go task("B")

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s\n", i, "anonymous")
			time.Sleep(time.Second)
		}
	}()
	time.Sleep(time.Second * 15)
}

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s\n", i, name)
		time.Sleep(time.Second)
	}
}
