### Closures

Um closure nada mais é do que uma função anônima declarada dentro do escopo de outra função e que possui acesso as
variáveis da função "pai". Exemplo utilizando a função soma da aula anterior:

```GO
func main() {
	result := func() int {
		return sum(1, 3, 45, 6, 34, 654, 654, 7645, 534, 543, 543, 543) * 2
	}()

	fmt.Println(result)
}
```
