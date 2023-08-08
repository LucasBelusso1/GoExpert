package main

import (
	"fmt"
)

func main() {
	result := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(result)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
