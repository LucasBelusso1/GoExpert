### FSNotify

A ideia aqui é, pensando em um cenário em que temos uma rotatividade de credenciais em um arquivo `.env`, por exemplo,
não queremos parar a aplicação para poder atualizar as nossas credenciais utilizando o arquivo que foi alterado.

Sendo assim, para simular um cenário parecido com o descrito acima, criaremos uma `main.go` com o seguinte conteúdo:"

```GO
type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBConfig

func main() {
	MarshalConfig("config.json")

	fmt.Println(config)
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		panic(err)
	}
}
```

No código acima, lemos um arquivo `config.json` as credenciais de conexão com o nosso banco de dados fictício e jogamos
a informação deste arquivo em uma variável `config` que é uma struct com as informações esperadas.
Em seguida criamos o arquivo `json` que conterá estas credenciais:

```JSON
{
	"db": "mysql",
	"host": "localhost",
	"user": "root",
	"password": "root"
}
```

Agora veja o código utilizando o FSNotify:

```GO
func main() {
	watcher, err := fsnotify.NewWatcher() // Cria um "watcher" ou "observador", que ficará observando o arquivo config.json

	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	MarshalConfig("config.json")

	done := make(chan bool) // Cria um canal para segurar a aplicação

	go func() { // Cria uma nova goroutine que monitora os eventos do watcher
		for {
			select {
			case event, ok := <-watcher.Events: // Recebe o evento
				if !ok {
					return
				}

				fmt.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write { // Verifica se é um evento de escrita
					MarshalConfig("config.json") // Atualiza o config com as informações atualizadas do config.json
					fmt.Println("modified file: ", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				fmt.Println(err)
			}
		}
	}()

	err = watcher.Add("config.json") // Adiciona o arquivo config.json para ser monitorado

	if err != nil {
		panic(err)
	}

	<-done
}

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		panic(err)
	}
}
```