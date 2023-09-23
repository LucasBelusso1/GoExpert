### Wait Groups

Veja o código abaixo:

```GO
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(3)

	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s\n", i, "anonymous")
			time.Sleep(time.Second)
		}
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s\n", i, name)
		time.Sleep(time.Second)
	}
	wg.Done()
}
```

Aqui contornamos o problema de espera que tínhamos na aula anterior utilizando `Wait groups`. Neste caso primeiramente
criamos uma instância de `sync.WaitGroup{}` e informamos a ele que vamos precisar esperar 3 goroutines serem executadas
para que o programa seja finalizado. Depois passamos o ponteiro do waitGroup para nossas funções que serão executadas
em paralelo e dentro de cada função, após o looping, informamos ao waitGroup que a tarefa foi finalizada. Ao final
das tarefas a thread principal é encerrada.
Caso definimos uma quantidade menor de processos dentro do `waitGroup.Add(3)` do que o necessário, algumas goroutines
serão executadas e outras não. O `waitGroup.Add(3)` funciona como um espécie de "crédito", e cada `Done()` que
chamamos, remove uma unidade deste crédito.