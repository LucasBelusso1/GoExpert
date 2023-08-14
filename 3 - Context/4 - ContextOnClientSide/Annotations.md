### Context no lado do Client

Utilizando o servidor que criamos na aula 3, com a regra de processamento definida em 5 segundos, podemos criar um
client http para realizar uma requisição e criar um contexto em cima deste client e definir um tempo para que o server
responda, do contrário a conexão será interrompida:

```GO
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}
```

Acima temos um código que chama o servidor que criamos e caso o servidor não responda em 3 segundos, a conexão é
interrompida, fazendo assim com que ambos o client e server parem de executar a thread.
Isso serve para mostrar que podemos controlar o contexto de execução tanto no lado do client quanto do lado do server.