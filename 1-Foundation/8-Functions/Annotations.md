### Funções

Em GO, as funções podem tanto receber quanto retornar múltiplos resultados.
Declaração de função:

```GO
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}

	return a + b, nil
}
```

A função acima requer 2 parâmetros do tipo int e retornará um inteiro (resultado da soma) e um erro, ou `nil` caso não
haja erro.
Chamando a função:

```GO
func main() {
	result, error := sum(10, 20) // output: 30 - nill.
	result2, error2 := sum(10, 60) // output: 0 - Error.

	if (error2 != nil) {
		fmt.Println("Ocorreu um erro", error2)
	}

	result3, _ := sum(10, 60) // Ignorando caso de erro.
}
```

Exemplo de função que retorna somente um parâmetro:

```GO
func sum(a, b int) int {
	return a+bz
}
```