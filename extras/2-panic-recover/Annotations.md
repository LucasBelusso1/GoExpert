### Panic Recover

A ideia aqui é, quando tivermos um "panic" em nosso projeto GO, fazer algo com este panic para compensar o pânico e
evitar com que o programa seja interrompido. Veja o exemplo abaixo:

```GO
func myPanic() {
	panic("Something went wrong")
}

func main() {
	myPanic()
}
```

No script acima, ao rodarmos um panico será retornado que irá interromper o programa. Agora veja com o panic recover:

```GO
func myPanic() {
	panic("Something went wrong")
}

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in main: ", r)
		}
	}()
	myPanic()
}
```

No script acima, pegamos o pânico gerado pela função `myPanic` e executamos uma ação em cima, que neste caso é imprimir
no console "recovered in main".

Também é possível tratar o pânico dependendo da mensagem retornada:

```GO
func panic1() {
	panic("panic1")
}
func panic2() {
	panic("panic2")
}

func main() {
	defer func() {
		r := recover()
		if r == "panic1" {
			fmt.Println("panic1 recovered")
		}

		if r == "panic2" {
			fmt.Println("panic2 recovered")
		}
	}()

	panic2() // Caso panic1() seja executado, a mensagem mudará de acordo com a validação que fizemos
}
```

Agora veja um exempo em um servidor HTTP:

```GO
func recoverMiddleware(next http.Handler) http.Handler { // Middleware para recuperação de pânicos.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Recoverd panic: %v\n", r)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	err := http.ListenAndServe(":8080", recoverMiddleware(mux))

	log.Println("Listening on port :8080")
	if err != nil {
		log.Fatalf("Could not listen on :8080: %v\n", err)
	}
}
```