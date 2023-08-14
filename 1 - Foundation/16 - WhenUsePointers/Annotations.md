### Quando usar ponteiros

Resumidamente, usa-se ponteiros quando é necessário alterar o valor de algo por referência dentro do escopo de outra
função ou processo. Exemplo:

```GO
func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}
```

No exemplo acima, ao informarmos os parâmetros para a função `soma()`, os valores dos parâmetros informados serão
modificados para 50 no lugar em que a função `soma()` foi executada:

```GO
func main() {
	number1 := 10
	number2 := 20
	soma(number1, number2)

	println(number1) // 50
	println(number2) // 50
}
```