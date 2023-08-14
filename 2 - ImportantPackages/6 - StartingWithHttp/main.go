package main

import (
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8082", nil)

	http.HandleFunc("/", searchCEP) // Passando uma função

	http.HandleFunc("/alternative", func(w http.ResponseWriter, r *http.Request) { // Passando uma função anônima
		w.Write([]byte("Hello World"))
	})

	if err != nil {
		fmt.Println("Error", err)
	}
}

func searchCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
