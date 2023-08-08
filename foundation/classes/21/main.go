package main

import (
	"fmt"

	"github.com/LucasBelusso1/GoExpert/foundation/classes/21/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Println("Resultado:", soma)

	fmt.Println(matematica.Number)
	// fmt.Println(matematica.number) // Resulta em erro

	car := matematica.Car{Something: "qualquercoisa"}

	fmt.Println(car)

	car.Run()
	// car.run() // Resulta em erro

}
