package main

import (
	"fmt"

	"github.com/LucasBelusso1/GoExpert/3-ExportingObjects/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}
