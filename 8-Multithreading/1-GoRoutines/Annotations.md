### Inciando com Go Routines.

Veja o código abaixo:

```GO
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s\n", i, name)
		time.Sleep(time.Second)
	}
}
```
No caso acima, temos uma função que inicia uma operação que imprime uma string a cada 1 segundo.

```GO
task("A")
task("B")
```

Ao chamar a função do jeito que está acima, a task "A" rodará até o final e somente depois a task "B" começará a ser
executada.

```GO
go task("A")
go task("B")
```

No caso acima, dizemos para o GO criar 2 novas threads usando a palavra reservada `go` antes da chamada da função,
entretanto o código acima não executará nada, pois temos que considerar o seguinte: Em go, existe a Thread principal
rodando, sendo assim, no cenário acima teríamos 3 threads, a principal e as 2 que foram criadas para executar as tasks,
sendo assim, caso a thread principal finalizar a execução e não der tempo de executar as goroutines, o programa será
encerrado.

```GO
go task("A")
go task("B")
time.Sleep(time.Second * 15)
```

Para resolver o problema da thread principal, podemos deixá-la dormindo por 15 segundos (pois sabemos que o nosso
código será executado em ~10s) até que as nossas goroutines terminem de exectar. Neste caso é uma solução paliativa
para o problema que será resolvida utilizando `Wait Groups`.

Também é possível executar uma goroutine em uma função anônima, desta forma:

```GO
go func() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s\n", i, "anonymous")
		time.Sleep(time.Second)
	}
}()
```