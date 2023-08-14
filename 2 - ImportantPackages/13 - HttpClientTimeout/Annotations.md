### HttpClient com Timeout

Com GO é possível determinar um tempo de timeout ao fazer uma requisição para algum recurso externo, passando
um parâmetro de timeout ao iniciar um httpClient, desta forma:

```GO
c := http.Client{Timeout: time.Second}
```

### Trabalhando com POST

Para fazer um Post com Go basta mudar a função que é chamada do httpClient criado para `Post()` e passar as informações
necessárias, que no caso são a url, o content type e um buffer de dados em formato de bytes. Veja o exemplo:

```GO
c := http.Client{}
jsonVar := bytes.NewBuffer([]byte(`{"name":"Lucas"}`))
resp, err := c.Post("https://google.com", "application/json", jsonVar)
```

### Customizando a request

É possível customizar a request que deseja fazer criando um `http.NewRequest`, assim será possível customizar tudo o que
diz respeito a requisição antes de submitá-la. E para realizar a requisição basta chamar o `Do` do httpClient que foi
criado:

```GO
c := http.Client{}
req, err := http.NewRequest("GET", "http://google.com", nil)
req.Header.Set("Accept", "application/json") // Customizando header
```

### Trabalhando com http usando contextos

Utilizando o pacote `context` é possível definir regras de execução para que aquele contexto de código seja interrompido
dependendo da regra definida.
Por exemplo, podemos definir um contexto para que a operação seja cancelada caso uma requisição demorar mais de 1
segundo para ser executada, veja o exemplo abaixo:

```GO
ctx := context.Background() // Criando contexto vazio
ctx, cancel := context.WithTimeout(ctx, time.Second) // Definindo regra de timeout de 1 segundo.

req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com.br", nil) // Requisição http passando o contexto.
```