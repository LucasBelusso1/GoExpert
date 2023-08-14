### Funções variádicas

Uma função variádica nada mais é do que uma função que pode receber inúmeros parâmetros. Entretanto, somente é possível
passar múltiplos valores do mesmo tipo. Por exemplo:

```GO
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
```

No exemplo acima a função recebe inúmeros valores inteiros e retorna a soma destes valores. É possível ter mais
parâmetros de outros tipos, entretanto apenas um deles pode ser variádico e este deve ser declarado por último.