### Generics

Generics é uma forma de conseguir passar mais de um tipo para uma função sem precisar usar a interface vazia.
Exemplo de como implementar Generics em uma função:

```GO
func Soma[T int | float64](numbers []T) T {
	var soma T
	for _, number := range numbers {
		soma += number
	}
	return soma
}
```

No exemplo acima, declaramos o nome da função como `Soma` e informamos dentro de colchetes `[]` que a função pode
receber ou um slice de int ou um slice de float64.

Também é possível definir uma interface que será um tipo com mais de um tipo, e substituir nos colchetes:

```GO
type Number interface {
	int | float64
}

func Soma[T Number](numbers []T) T {
	var soma T
	for _, number := range numbers {
		soma += number
	}
	return soma
}
```

Caso haja um tipo de seja criado com um tipo subjacente (`type AnyNumber int`), o exemplo acima irá quebrar. Para que o
GO reconheça que o tipo executado possui um tipo subjacente válido para realizar a operação, é necessário informar a
interface que pode haver essa possibilidade. Para fazer isso basta adicionar o caractere `~` a esqueda dos tipos,
da seguinte forma:

```GO
type Number interface {
	~int | ~float64
}
```
