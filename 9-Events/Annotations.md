### O que são eventos

Eventos são acontecimentos do passado, e a partir desse acontecimento é possível fazer com que o software tome alguma
ação.
Por exemplo, um evento de inserção de cliente (Evento: Cliente inserido). A partir deste evento podemos executar
diversas ações como disparo de e-mail, publicar uma mensagem em uma fila, notificar um usuário no slack...

Um evento é composto por alguns itens:

`Evento`: É a própria ação a ser executada, que pode ou não carregar dados consigo.
`Operações`: Operações que serão executadas durante ou após a execução deste evento.
`Gerenciador`: O gerenciador de eventos que é o coração do sistema de eventos, é ele que vai gerenciar qual evento foi
executado e quais as operações, e também é ele quem vai dar o **dispatch/fire** (disparar) o evento.

### Criação das interfaces:

Criamos então um projeto e nele criamos a pasta `/pkg/events` com o arquivo `interfaces.go`. Dentro deste arquivos
inserimos o seguinte código:

```GO
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
```

Neste caso, criamos algumas interfaces, dentre elas:

`EventInterface` que vai conter um nome, a data e hora da execução e o payload.
`EventHandlerInterface` que vai representar a operação a ser executada no evento.
`EventDispatcherInterface` que vai gerenciar todos os eventos, tendo as funções:
- `Register`: Para guardar um ou mais eventos/operações.
- `Dispatch`: Para executar o evento.
- `Remove`: Para remover um evento.
- `Has`: Para verificar se um evento existe.
- `Clear`: Para limpar todos os eventos registrados.

### Criando event dispatcher

Veja o código abaixo:

```GO
var ErrHandlerAlreadyRegistered = errors.New("hanlder already registered")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	_, ok := ed.handlers[eventName]
	if ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
```

Aqui então criamos uma `struct` que vai conter a propriedade **handlers** que será um map chave <-> valor em que a chave
será o nome do evento e o valor sera uma slice de eventos. Em seguida criamos a função `NewEventDispatcher()` que apenas
cria uma nova instância da nossa struct e criamos a função `Register()` atrbuindo-a ao nosso `eventDispatcher` que
recebe o nome do evento e o que será executado. Na função `Register` validamos se aquele handler já existe dentro do
evento, caso exista retornamos um erro, do contrário adicionamos a slice.

### Criando suite de testes

Neste caso, para fazer o setup de testes, precisamos criar alguns events e eventHandlers fictícios, e para isso, além
de criar as structs e suas funções, utilizamos o pacote `github.com/stretchr/testify/suite` para criar uma suite de
testes que será utilizada em todas as funções:

```GO
type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
```

**OBS.:** A função SetupTest roda para cada função de teste que temos.

### Testando register

Para testar o register, utilizamos o suite para pegar os dados do setup e utilizá-los:

```GO
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Error(err)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
}
```

### Criando e testando o Clear

Para criar a função `Clear()` é muito simples, basta reatribuir ao handler um map de string e slice de EventHandlers:

```GO
func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]EventHandlerInterface)
}
```

E para testá-lo fazemos da seguinte forma:

```GO
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Nil(err)
	suite.Empty(suite.eventDispatcher.handlers)
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}
```

### Criando e testando o Has

Para criar o Has também é bem simples e semelhante ao `Register()`:

```GO
func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	_, ok := ed.handlers[eventName]
	if ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}
```

E para criar os testes:

```GO
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	hasEvent := suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(hasEvent)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	hasEvent = suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2)
	suite.True(hasEvent)

	hasEvent = suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3)
	suite.False(hasEvent)

	hasEvent = suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler3)
	suite.False(hasEvent)
}
```

### Criando e testando o Dispatch

Para criar o Dispatch, fazemos da seguinte forma:

```GO
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	handlers, ok := ed.handlers[event.GetName()]

	if ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}

	return nil
}
```

Primeiramente verificamos se há handlers disponíveis para aquele evento, em seguida passamos por cada handler o
executando chamando o método `handle`.

Agora para criar o teste, precisamos criar primeiramente um mock de um `EventHandler` utilizando o pacote
`github.com/stretchr/testify/mock`, desta forma:

```GO
type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}
```

Agora no teste escrevemos da seguinte forma:

```GO
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &event) // Ao chamar o `Handle` do mock, passamos o event como parâmetro
	dispacher.Register(event.GetName(), eh)
	dispacher.Dispatch(&event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1) // Verificamos se a função Handle do mock foi chamada 1 vez.
}
```

### Slice of a slice

Para entender a remoção de eventos, precisamos antes revisitar o conceito de slices.
Para slices, podemos definir um "range" para ser extraido do slice. Veja o exemplo:

```GO
slice := []string{"banana", "maçã", "laranja", "uva"}
slice = slice[1:2]
fmt.Println(slice) // [maçã]
```
No exemplo acima, informamos ao GO que queremos, a partir da primeira posição, pegar até o item 2, neste caso, a
primeira posição é "banana", e o item 2 é "maçã", então o que vai ser impresso é [maçã].
Veja outro exemplo:

```GO
slice := []string{"banana", "maçã", "laranja", "uva"}
slice = slice[1:]
fmt.Println(slice) // ["maçã", "laranja", "uva"]
```

No caso acima, informamos ao GO que queremos a partir da posição 1, todos os elementos.

```GO
slice := []string{"banana", "maçã", "laranja", "uva"}
slice = slice[:3]
fmt.Println(slice)
```
No caso acima, informamos ao GO que queremos da posição zero até a posição 3 (incluindo).

### Criando e testando Remove

Entendendo o conceito anterior de slices, agora vamos aplicá-la em nossa função `remove()`, desta forma:

```GO
func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	_, ok := ed.handlers[eventName]
	if ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return ErrHandlerNotFound
}
```

Para testá-la, segue o modelo dos demais testes:

```GO
func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	// Adiciona Evento 1 Handler 1
	err := dispacher.Register(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))

	// Adiciona Evento 1 Handler 2
	err = dispacher.Register(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(2, len(dispacher.handlers[event.GetName()]))

	// Adiciona Evento 2 Handler 3
	err = dispacher.Register(event2.GetName(), &handler3)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event2.GetName()]))

	// Remove Handler 1 do Evento 1
	dispacher.Remove(event.GetName(), &handler)
	suite.Nil(err)
	suite.Equal(1, len(dispacher.handlers[event.GetName()]))
	suite.Equal(&handler2, dispacher.handlers[event.GetName()][0])

	// Remove novamente handler 1 (erro)
	dispacher.Remove(event.GetName(), &handler)
	suite.Error(ErrHandlerNotFound)

	// Remove Handler 2 do Evento 1
	dispacher.Remove(event.GetName(), &handler2)
	suite.Nil(err)
	suite.Equal(0, len(dispacher.handlers[event.GetName()]))

	// Remove Handler 3 do Evento 2
	dispacher.Remove(event2.GetName(), &handler3)
	suite.Nil(err)
	suite.Equal(0, len(dispacher.handlers[event2.GetName()]))
}
```

### Adicionando Go routines

A ideia aqui é transformar a função `dispatch` de modo que os handlers sejam executados de forma paralela. Sendo assim
precisamos modificar a função da seguinte forma:

```GO
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	handlers, ok := ed.handlers[event.GetName()]

	if ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}

	return nil
}
```

No código acima, definimos que o handler será executado em uma outra thread a partir da palavra reservada `go` e
definimos um waitgroup que esperará para que a **goroutine** seja executada. Para cada iteração do for, adicionamos mais
um "crédito" em nosso waitgroup.

Por conta da goroutine, devemos alterar a função `Handle` para que receba um waitgroup, isso tanto na chamada da função
quanto na interface e nos testes.

```GO
// Interface
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}
```
```GO
// Testes
func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &event)

	dispacher.Register(event.GetName(), eh)
	dispacher.Register(event.GetName(), eh2)
	dispacher.Dispatch(&event)
	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
```

### Instalando RabbitMQ

Para instalar, basta criarmos o arquivo `docker-compose.yaml` com os seguintes dados:

```YAML
version: '3'

services:
  rabbitmq:
    image: rabbitmq:3.8.16-management
    container_name: rabbitmq
    hostname: rabbitmq
    ports:
      - "5672:5672"
      - "15672:15672"
      - "15692:15692"
    environment:
      - RABBITMQ_DEFAULT_USER=guest
      - RABBITMQ_DEFAULT_PASS=guest
      - RABBITMQ_DEFAULT_VHOST=/
```

Em seguida, executamos `docker compose up -d`.
É possível acessar a área de administrador pela URL `http://localhost:15672`

[Como o RabbitMQ funciona.](https://tryrabbitmq.com/)

### Consumindo mensagens.

Para criar um consumidor, utilizamos o pacote `github.com/rabbitmq/amqp091-go`, desta forma:

```GO
func OpenChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume(
		"myqueue",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}
```

No código acima, criamos um pacote `rabbitmq` com as funções `OpenChannel()` e `Consume()`, nos quais:

`OpenChannel()`: Abre uma conexão com o RabbitMQ e cria um canal.
`Consume()`: A partir do canal recebido por parâmetro, consome as mensagens que estão na fila **myqueue** e envia para
o channel do GO.

Agora criamos dentro de `cmd/consumer` o arquivo `main.go` e nele definimos o seguinte:

```GO
func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs)

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
```

No código acima, abrimos a conexão com o RabbitMQ e criamos o canal, em seguida criamos um channel do tipo
`amqp.Delivery` e dizemos para que as mensagens sejam consumidas paralelamente. Em seguida capturamos a informação
destas mensagens consumidas pelo canal, jogamos no stdout e executamos um `msg.Ack(false)` que informa para o RabbitMQ
que a mensagem foi consumida e não precisa retornar a fila.

### Produzindo e consumindo mensagens.

Para produzir uma mensagem, dentro do pacote `rabbitmq` que criamos para as funções de `OpenChannel()` e `Consume()`
escrevemos o seguinte código:

```GO
func Publish(ch *amqp.Channel, body string, exchange string) error {
	err := ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
```

Na função acima, recebemos o canal, o conteúdo da mensagem e qual será a exchange utilizada. A partir dela, criamos
dentro de `cmd/producer` um outro arquivo main para produzir a mensagem, desta maneira:

```GO
func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()
	rabbitmq.Publish(ch, "Hello World!", "amq.direct")
}
```

**OBS.:** Para que funcione corretamente, será necessário acessar `http://localhost:15672` e fazer um bind da queue
que criamos com a exchange default `amq.direct`.