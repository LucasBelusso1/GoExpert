package matematica

import "fmt"

var number int = 10 // Acessível somente dentro do pacote
var Number int = 10 // Acessível fora do pacote

type car struct { // Acessível somente dentro do pacote
	something string // Acessível somente dentro do pacote
}

type Car struct { // Acessível fora do pacote
	Something string // Acessível fora do pacote
}

func (c Car) Run() {
	fmt.Println("carro está andando!")
}

func (c Car) run() {
	fmt.Println("carro está andando!")
}

func Soma[T int | float64](a, b T) T {
	return a + b
}
