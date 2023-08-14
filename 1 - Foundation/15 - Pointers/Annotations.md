### Ponteiros

`&` -> Quando utilizado o caractere de "E comercial", o GO buscará o endereço da memória daquela variável.
Exemplo:

```GO
func main() {
	number := 10
	// Resultará em algo assim: 0xc00004a768, que é o endereço da memória em que este valor está armazenado.
	println(&number)
}
```

`*` -> Quando usado o asterísco, o GO tentará buscar o valor de um endereço de memória.
Exemplo:

```GO
func main() {
	number := 10
	pointer := &number

	// Resultará em 10, pois ele busca o valor no endereço fornecido.
	println(*pointer)
}
```