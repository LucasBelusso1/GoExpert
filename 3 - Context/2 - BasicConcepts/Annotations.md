### Entendendo conceitos básicos.

Como mencionado na aula anterior, é possível utilizar regras para que algo seja executado, seja tempo, valor...
Veja o exemplo:

```GO
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled! Timeout reached.")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked!")
	}
}
```

No caso acima, criamos um contexto e definimos uma regra para que ele seja cancelado em 3 segundos e passamos este
contexto para a função `bookHotel()`, que dependendo do status do contexto ou do tempo decorrido executa diferentes
tarefas.
Neste cenário o hotel nunca será reservado pois o contexto finaliza a execução antes dos 5 segundos definidos no
`select`.