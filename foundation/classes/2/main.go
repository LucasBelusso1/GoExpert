package main

// Declaração em escopo global

const helloWorld = "Hello World"

var boolExample bool

var (
	verifier bool    // Default false
	number   int     // Default 0
	text     string  // Default ""
	fraction float64 // Default +0.000000e+000
)

var (
	verifier2 bool    = true
	number2   int     = 90
	text2     string  = "Hello World!"
	fraction2 float64 = 5.546
)

func main() {
	// Constante
	println(helloWorld)

	// Variável
	println(boolExample)

	// Múltiplas variáveis
	println(verifier)
	println(number)
	println(text)
	println(helloWorld)

	// Múltiplas variáveis com valor atribuido
	println(verifier2)
	println(number2)
	println(text2)
	println(fraction2)

	duckTypeVariable := "Hello World!"

	println(duckTypeVariable)
}
