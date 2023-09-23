### Comando.

O GO possui o próprio comando para identificar se o seu código esta sofrendo com race conditions, basta executar:

```SHELL
go run -race {{arquivo}}.go
```

### Mutex

Utilizando `Mutex` do pacote `sync`, é possível informar ao GO para que ele bloqueie o acesso a uma variável que está
sofrendo de uma race condition desta forma:

```GO
m := sync.Mutex{}
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	number++
	m.Unlock()
	w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
})
```

### Atomic

Utilizando o pacote `atomic`, é possível simplificar um pouco o código sem ter que ficar dando `Lock()` e `Unlock()` em
cada variável que precisarmos tratar uma race condition, desta forma:

```GO
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&number, 1)
	w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
})
```