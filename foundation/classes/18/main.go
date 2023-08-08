package main

import (
	"fmt"
)

func main() {
	var number int = 10
	var text string = "Hello World!"

	PrintTypeAndValue(number)
	PrintTypeAndValue(text)
}

func PrintTypeAndValue(any interface{}) {
	fmt.Printf("O valor da variável é %v e o tipo da variável é %T\n", any, any)
}
