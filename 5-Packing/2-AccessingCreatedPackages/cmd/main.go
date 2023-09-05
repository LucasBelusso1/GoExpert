package main

import (
	"fmt"

	"github.com/LucasBelusso1/GoExpert/2-AccessingCreatedPackages/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
