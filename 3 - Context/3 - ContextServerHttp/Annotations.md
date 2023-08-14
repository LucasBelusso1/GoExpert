### Context utilizando server HTTP

Em uma requisição http, um contexto é enviado para o `http.Request()`, sendo assim, é possível capturar ações do client
que foram executadas em cima deste contexto. Por exemplo, caso o client cancele a requisição por algum motivo,
utilizando o `ctx.Done()` é possível interromper a execução daquela tarefa para que não seja consumido processamento
sem necessidade.

```GO
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada!")
	defer log.Println("Request finalizada!")

	select {
	case <-time.After(time.Second * 5):
		// Imprime no CLI (stdout)
		log.Println("Request processada com sucesso!")

		// Imprime no browser
		w.Write([]byte("Request processada com sucesso!"))
	case <-ctx.Done():
		// Imprime no CLI (stdout)
		log.Println("Request cancelada pelo cliente.")
	}
}
```

No exemplo acima criamos um servidor que após passados 5 segundos retorna uma resposta de sucesso, e caso o contexto
seja cancelado por alguma razão o processo é finalizado e a thread é fechada.