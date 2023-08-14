### Context WithValue

É possível passar e resgatar valores a partir de chaves para um contexto, apesar de não ser muito utilizado é uma
possibilidade dependendo do cenário.

```GO
func main() {
	ctx := context.WithValue(context.Background(), "token", "senha")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")

	fmt.Println(token)
}
```

No exemplo acima é criado um contexto e dentro dele definimos a chave "token" e o seu valor "senha", e na função
`bookHotel()` resgatamos este valor.

**IMPORTANTE:** Sempre que estamos passando contextos para funções, o parâmetro do contexto, por convenção, deve sempre
ser o primeiro.