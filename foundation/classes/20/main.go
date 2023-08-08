package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](numbers []T) T {
	var soma T
	for _, number := range numbers {
		soma += number
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	sliceOfInt := []int{1000, 2000, 3000}
	sliceOfFloat64 := []float64{100.20, 2000.3, 300.0}

	sliceOfMyNumber := []MyNumber{1000, 2000, 3000}
	println(Soma(sliceOfInt))
	println(Soma(sliceOfFloat64))
	println(Soma(sliceOfMyNumber))
	println(Compara(10, 10))
}
