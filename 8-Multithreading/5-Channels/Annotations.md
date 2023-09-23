### Channels

Para trabalhar com channels primeiro precisamos entender como ele funciona. O channel em GO é como se fosse um canal
que trabalha em dois estados, cheio ou vazio. Quando vazio, a thread que solicitou o dado fica esperando até que o dado
seja disponibilizado por outra thread. Uma vez que o canal esteja cheio, não é possível inserir mais informações.

veja um exemplo em código:

```GO
channel := make(chan string) // Empty channel

go func() {
	channel <- "Hello World!" // Full channel
}()

msg := <-channel // Channel gets empty again
fmt.Println(msg)
```

### Forever (deadlock)

O deadlock acontece quando temos uma thread que espera um dado de um canal que nunca é populado. Veja o exemplo:

```GO
forever := make(chan bool)

<-forever
```

No caso acima, criamos um canal e dizemos para a thread principal buscar um valor deste canal, porém este valor nunca é
instanciado.
Para resolver este problema, a única solução possível é, em outra thread, popular a informação deste canal, não é
possível popular o canal na mesma thread em que ele foi criado.

```GO
go func() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	forever <- true
}()

<-forever
```

### Iterando com range

Veja o exemplo abaixo:

```GO
func main() {
	channel := make(chan int)

	go publish(channel)
	reader(channel)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
```

Neste caso criamos um channel e chamamos a função `publish` que faz um looping de 10 iterações adicionando valores ao
channel, e temos a função `reader` que lê a informação que é inserida no channel. O que acontecerá neste cenário é que
o `publish` irá inserir os 10 dígitos dentro do channel e cada vez que o dígito é inserido o `reader` buscará esta
informação e irá imprimir no console. Entretanto, o que o `reader` está fazendo é aguardar infinitamente por uma
informação, e sempre que a informação é inserida ele a captura, porém como o publish se limita a 10 iterações, será
exibido um erro de deadlock no final da execução pois o `reader` parará de receber informações e o channel ficará
aberto.

Para resolver isso, ao final da iteração do nosso loop, informar que o channel não receberá mais nenhum valor,
fechando-o, desta forma:

```GO
func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
```

### Range com WaitGroups

Veja o código abaixo:

```GO
func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go publish(channel, &waitGroup)
	go reader(channel)
	waitGroup.Wait()
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publish(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}
```

No exemplo acima, executamos tanto o `publish` quanto o `reader` em threads diferentes, e para que isso funcione,
precisamos informar ao GO que será necessário aguardar até que as threads terminem de executar. Para isso, criamos um
WaitGroup e ao final do looping da função `publish`, fechamos o canal e também terminamos o WaitGroup.

### Directions

É possível definir uma "direção" para os canais, fazendo com que o canal somente possa ser lido ou somente escrito.
Veja o exemplo:

```GO
func main() {
	channel := make(chan string)

	go recieve("Hello", channel)
	read(channel)
}

func recieve(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}
```

No caso a função `recieve` apenas pode ler as informações do channel enquanto que `read` pode apenas inserir
informações no canal.
É recomendado que sempre utilize-se directions para trabalhar com channels para garantir que não haverá erros.

### Criando LoadBalancer

Veja o código abaixo:

```GO
func main() {
	channel := make(chan int)

	// Initialize workers
	for i := 0; i < 1; i++ {
		go worker(i, channel)
	}

	// Request simulation
	for i := 0; i < 100; i++ {
		channel <- i
	}
}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
```

No código acima, criamos uma função chamada `worker` consome uma informação do channel (se estiver disponível) e aguarda
1 segundo, simulando o processamento de uma requisição, e depois criamos um `for` que popula esse channel simulando as
requests.
Caso adicionarmos mais workers para trabalhar, as requisições serão processadas mais rapidamente em paralelo, assim como
em um load balancer.

### Trabalhando com Select

Veja o exemplo:

```GO
type Message struct {
	id  int64
	msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var counter int64 = 0

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			c1 <- Message{counter, "Hello from RabbitMQ"}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			c2 <- Message{counter, "Hello from Kafka"}
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("Received from Kafka, ID: %d - %s\n", msg.id, msg.msg)
		case msg := <-c2:
			fmt.Printf("Redceived from RabbitMQ, ID: %d - %s\n", msg.id, msg.msg)
		case <-time.After(time.Second * 4):
			fmt.Println("Timeout")
		}
	}
}
```

No cenário acima, criamos um `select` dentro de um looping infinito o qual fica monitorando para saber se algum channel
recebeu alguma informação, e caso não receba dentro de 4 segundos, retorna um timeout.
O `select` geralmente é utilizado para verificar o retorno de mais de um channel, e é executado sempre o primeiro
channel que retornar informações.

### Buffer

O buffer permite com que seja adicionado mais do que um valor dentro de um canal, veja o exemplo sem utilizar buffer:

```GO
ch1 := make(chan string)

ch1 <- "hello"
ch1 <- "world"

fmt.Println(<-ch1)
fmt.Println(<-ch1)
```

Este código resultará em erro de deadlock, pois estaremos tentando adicionar mais valores do que o canal permite. Agora
veja o exemplo com buffer:

```GO
ch1 := make(chan string, 2)

ch1 <- "hello"
ch1 <- "world"

fmt.Println(<-ch1)
fmt.Println(<-ch1)
```
Acima, indicamos para o make que o canal terá um buffer de 2 valores, assim sendo possível adicionar os valores de
"hello world" para dentro do canal. Entretanto o uso de buffer geralmente não é recomendado e requer bastante atenção
caso seja necessário utilizar.