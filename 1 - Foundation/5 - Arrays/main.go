package main

import "fmt"

const a = "Hello World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40 // Resulta em erro.

	fmt.Println("A última posição do array é", len(myArray)-1)
	fmt.Println("O último valor do array é", myArray[len(myArray)-1])

	fmt.Println("As posições do meu array e seus respectivos valores são:")
	for index, value := range myArray {
		fmt.Println("O valor do índice é", index, "e o valor é", value)
	}
}
