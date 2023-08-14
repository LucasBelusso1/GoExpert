### Interfaces vazias

A interface vazia `interface{}` é uma interface que está presente em todas as structs, variáveis, tipos no geral. Sendo
assim, é possível fazer com que uma função, por exemplo, possa receber mais de um tipo em uma mesma variável. Veja o
exemplo:

```GO
func main() {
	var number int = 10
	var text string = "Hello World!"

	PrintTypeAndValue(number)
	PrintTypeAndValue(text)
}

func PrintTypeAndValue(any interface{}) {
	fmt.Printf("O valor da variável é %v e o tipo da variável é %T\n", any, any)
}
```

No exemplo acima, o parâmetro `any` recebe tanto uma string quanto um int e não retorna erro. Entretanto nestes casos
é necessário realizar uma série de validações para não violar a tipagem do GO.