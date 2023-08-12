package main

import (
	"context"
	"io"
	"net/http"
)

func main() {
	ctx := context.Background()
	// Cancela o contexto ao executar cancel() ou ao final do tempo
	// ctx, cancel := context.WithTimeout(ctx, time.Second)

	// Cancela o contexto apenas ao executar a função cancel()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com.br", nil)

	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
