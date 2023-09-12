### Iniciando com testes automatizados

Por convenção, para realizar testes em GO, cria-se um arquivo com o nome do arquivo que você está querendo testar com
o sufixo `_test`. Por exemplo, se estou querendo testar o arquivo `tax.go`, será criado o arquivo `tax_test.go` para
comportar todos os testes.

No arquivo `tax_test.go`, supondo que queremos testar a função `CalculateTax()`, criamos então a função
`TestCalculateTax()`, sempre utilizando o prefixo `Test`, na qual definimos o que vamos mandar pra função e o que
esperamos de retorno, e utilizando o pacote `testing` e as verificações necessárias, informamos ao GO em caso de erro
através da função `t.Errorf()` que o teste falhou, passando o resultado e o resultado esperado.

Para testar múltiplos valores para uma mesma função, basta mudar a lógica da própria função para trabalhar com multiplas
comparações. Um mesmo teste pode fazer inúmeros testes.

### Code coverage

Para verificar qual é o percentual do seu código que está sendo testado, é possível rodar o comando
`go test -coverprofile=coverage.out`, que informará qual é o percentual e gerará o arquivo `coverage.out` que informará
quais linhas estão sendo testadas e quais não passaram pelo teste. Para ver este mesmo dado em html basta rodar o
seguinte comando `go tool cover -html=coverage.out`, após rodar o comando anterior.

### Benchmark

Para fazer um benchmark da nossa função, primeiro criamos uma função de Benchmark dentro do arquivo `tax_test.go`,
dentro dele receberemos `b *testing.B` que é a parte do pacote de testes responsável pelo benchmarking:

```GO
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
```

Após criar a função, executaremos o seguinte comando:

```SHELL
go test -bench=.
```
Ou então, para executarmos somente o benchmark, utilizamos:

```SHELL
go test -bench=. -run=^#
```

Utilizando o recurso de benchmark é possível gerar funções iguais com comportamentos diferentes e testá-las para
verificar qual das duas é mais performática.

### Fuzzing

O Fuzzing é uma estratégia dentro dos testes que força diversos valores diferentes nos parâmetros da função para que não
seja necessário escrever todos os valores possíveis.

Exemplo de utilização do fuzzing:

```GO
func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)

		if amount <= 0 && result != 0 {
			t.Errorf("Expected 0 but got %f", result)
		}

		if amount > 20000 && result != 20.0 {
			t.Errorf("Expected 20 but got %f", result)
		}
	})
}
```

No caso acima, declaramos alguns valores para serem utilizados como base e adicionamos ao fuzzing. Em seguida, chamamos
a função `Fuzz` que recebe como parâmetro um ponteiro para testing.T e os parâmetros da função que estamos querendo
testar e inserimos as assertions que desejamos fazer.

Para rodar o teste utilizando fuzzing, utilizamos o seguinte comando:

```SHELL
go test -fuzz=. -fuzztime=5s -run=^#
```
Em que:

- `-fuzz=.` indica para o GO será utilizado fuzzing.
- `-fuzztime=5s` que é o tempo em que o teste irá rodar enviando valores aleatórios para a função.
- `-run=^#` para evitar rodar testes sem fuzzing.

Ao rodar um teste utilizando fuzz, em casos de erro o sistema gerará uma pasta chamada **testData/fuzz/fuzz{{Função}}**
com arquivos contendo qual foi o valor utilizado que resultou em erro. Ao corrigir o erro, é possível rodar este mesmo
teste com o comando:

```GO
go test -run=FuzzCalculateTax/{{HASH_DO_ARQUIVO}}
```