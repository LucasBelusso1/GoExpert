package main

import "fmt"

func main() {
	// Memória -> Endreço -> Valor
	number := 10
	fmt.Println("O valor de number é", number)

	var pointer *int = &number

	fmt.Println("O endereço de memória de number é", pointer)
	fmt.Println("O valor do endereço de memória de number é", *pointer)

	*pointer = 30

	fmt.Println("O valor de number é", number)
}
