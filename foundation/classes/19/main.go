package main

import (
	"fmt"
)

func main() {
	var minhaVar interface{} = "Wesley Willians"
	// Caso não passar o .(string), como o GO não sabe o tipo, irá imprimir o endereço da memória
	println(minhaVar)
	println(minhaVar.(string))

	// Convertendo e verificando se a conversão funcionou
	// Caso funcione, "res" vai receber o valor e "ok" será true, do contrário "res" receberá 0 e "ok" será false
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v", res, ok)

	// Neste caso não está sendo verificado o "ok" e a conversão resultará em erro.
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
