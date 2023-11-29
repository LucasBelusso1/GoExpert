### Graceful shutdown

O objetivo aqui é fazer com que um processo seja interrompido de forma que não prejudique o que está sendo processado.

Para exemplificar, primeiramente criamos um servidor com um sleep de 4 segundos na rota `"/"` dentro da main.go:

```GO
package main

import (
	"net/http"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World"))
	})

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
```
Em seguida, executamos este servidor (rodando o comando `go run main.go`) e fazemos uma chamada via `curl` para a porta
definida no script (neste caso 8080), e antes que a aplicação consiga terminar de processar a requisição, ou seja,
dentro do tempo do sleep, no terminal em que foi executado o nosso servidor, vamos executar um `ctrl + c` para matar
o processo.
Veja que neste caso o servidor não conseguiu terminar de processar a requição, retornando uma resposta me branco para
o client.

Agora para implementar o Graceful Shutdown fazemos modificamos o script da seguinte maneira:

```GO
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Hello World"))
	})

	go func() { // Separamos o servidor em uma outra thread.
		err := server.ListenAndServe()

		if err != nil && http.ErrServerClosed != err { // Verificamos se o erro apresentado não é um desligamento do servidor.
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1) // Criamos um canal para escutar um sinal do sistema operacional
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // Notifica o canal caso haja alguma interrupção (ctrl + c) ou sinal de interrupção do OS.
	<-stop // Escuta o canal até que o mesmo receba alguma notificação.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Define um contexto de 5 segundos que será executado caso alguma notificação entre no canal
	defer cancel()
	fmt.Println("Shutting down server...")

	err := server.Shutdown(ctx) // Desliga o servidor após 5 segundos.
	if err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped.")
}
```

No script acima, caso seguirmos os mesmos passos de antes, agora ao executar o `ctrl + c`, o server demorará 5 segundos,
processará o hello world solicitado e somente depois fechará o servidor.