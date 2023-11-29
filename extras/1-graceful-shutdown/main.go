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

	go func() {
		err := server.ListenAndServe()

		if err != nil && http.ErrServerClosed != err {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Shutting down server...")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped.")
}
