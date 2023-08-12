# Pacotes importantes

### 1 - Manipulação de arquivos

Em Go é possível fazer algumas operações com arquivos utilizando o pacote `os`. Nele é possível chamar a função
`Create()` para criar um arquivo (que retorna uma instância no `os.File`), `ReadFile()` ou `Open` para ler o conteúdo de
um arquivo e `Remove()` para deletar o arquivo, entre outras funções...

Com a instância do `os.File` retornada ao criar ou ler um arquivo, também é possível executar algumas funções para
manipular o arquivo em questão, como `os.File.Write()` que recebe um slice de bytes para inserir no arquivo,
`os.File.WriteString()` para inserir somente strings, entre outros recursos.

Também é possível ler o arquivo gradualmente (caso ele seja grande demais e aloque muito espaço na memória), utilizando
o pacote `bufio`, criando um novo Reader com `bufio.NewReader()` passando o arquivo em questão e passando buffer que
deverá ser respeitado ao ler o arquivo a função `Read()`.

### 2 - Realizando chamadas http

É possível realizar chamadas http utilizando o pacote `net/http` para fazer as operações, como GET, POST, DELETE,
PATCH, PUT...entre outras funções.
Para fazer a leitura da resposta de alguma requisição, utiliza-se o pacote `io`, com a função `ReadAll` passando o
response body.

### 3 - Defer

O `defer` statement é uma palavra reservada que quando utilizada, faz com que o trecho de código execute ao final da
execução (geralmente) de uma função.

### 4 - Trabalhando com JSON

Em GO, fazemos o mapeamento de JSON's para `Structs`, utilizando o pacote `json`. Ou seja, para transformar um JSON
em um "objeto" no GO, precisamos ter uma `struct` com as propriedades daquele json, somente após ter a struct,
executamos a função `Unmarshal()` do pacote `json` passando uma struct vazia como parâmetro para que as informações
sejam prenchidas.

Já para transformar a nossa `struct` em um JSON, é necessário criar um objeto daquela struct e então chamar a função
`Marshal()` que retornará um binário que precisa ser convertido para string para obter o JSON.

Também dentro do pacote `json` é possível definir encoders que enviarão o JSON para algum output definido.

### 5 - Busca CEP

Neste projeto utilizamos muitas das coisas que vimos para gerar um script que busca o CEP por meio de um valor fornecido
via parâmetro no terminal.
Para fazer a chamada HTTP foi utilizado o pacote `net/http` com a função `Get()`.
Para buscar o parâmetro do terminal foi utilizado o pacote `os` pegando a propriedade `Args`.
Utilizou-se o pacote `io` com a função `ReadAll()` para ler o response body.
Utilizou-se o pacote `json` para encodar e decodar o JSON retornado pela API.

### 6 - Iniciando com HTTP

Para criar um servidor em GO, basta utilizar a função `ListenAndServe()` to pacote `net/http`, a partir dela,
especifica-se uma porta e basta rodar o script que o servidor estará rodando.

Para criar rotas, utiliza-se a função `HandleFunc()` também to pacote `net/http` passando uma função que executará
aquela tarefa como parâmetro.

### 7 - Manipulando Headers

É possível retornar diferentes respostas HTTP dependendo da necessidade, utilizando o `ResponseWriter` com a função
`WriteHeader()` passando alguma das propriedades disponíveis no pacote `net/http`.
Assim como também é possível definir headers para serem retornados utilizando `.Header().Set()` ou então pegar
parâmetros enviados na URL utilizando o `URL.Query().Get()` do `Request`.

### 8 - Criando função buscaCEP

Nesta aula apenas mesclamos o que aprendemos na aula 6 e 7, iniciando um servidor e criando uma função que, a partir
de uma chamada para nosso `localhost:{port}`, conseguimos buscar um CEP do viacep.

### 9 - Finalizando resposta para o servidor HTTP

Nesta aula foi utilizado o pacote `json` pegar a struct gerada ao consultar o viacep e retornar em formato JSON
para o client.

### 10 - Server mux

Serve mux é uma forma de instanciar servidores com mais flexibilidade, podendo passar structs inteiras para atuar
como handlers e também atribuir rotas a diferentes portas.
É possível criar um mux através da função `NewServeMux()` do pacote `http`, passando-o como parâmetro para a função
`ListenAndServe()` do pacote `http`.

### 11 - File server

É possível fazer um servidor arquivo utilizando mux, criando um `FileServer()` a partir do pacote `http` e passando ele
para um server mux no parâmetro da função `Handle`. A função `FileServer()` recebe um parâmetro, que neste caso passamos
um `http.Dir({caminho_da_pasta})`

### 12 - Iniciando com templates

É possível fazer um servidor arquivo utilizando mux, criando um `FileServer()` a partir do pacote `http` e passando ele
para um server mux no parâmetro da função `Handle`. A função `FileServer()` recebe um parâmetro, que neste caso passamos
um `http.Dir({caminho_da_pasta})`

Dentro da pasta **12 - Templates** há alguns exemplos de como utilizar os templates de diferentes formas.

### 13.1 - HttpClient com Timeout

Com GO é possível determinar um tempo de timeout ao fazer uma requisição para algum recurso externo, passando
um parâmetro de timeout ao iniciar um httpClient, desta forma:

```GO
c := http.Client{Timeout: time.Second}
```

### 13.2 - Trabalhando com POST

Para fazer um Post com Go basta mudar a função que é chamada do httpClient criado para `Post()` e passar as informações
necessárias, que no caso são a url, o content type e um buffer de dados em formato de bytes. Veja o exemplo:

```GO
c := http.Client{}
jsonVar := bytes.NewBuffer([]byte(`{"name":"Lucas"}`))
resp, err := c.Post("https://google.com", "application/json", jsonVar)
```

### 13.3 - Customizando a request

É possível customizar a request que deseja fazer criando um `http.NewRequest`, assim será possível customizar tudo o que
diz respeito a requisição antes de submitá-la. E para realizar a requisição basta chamar o `Do` do httpClient que foi
criado:

```GO
c := http.Client{}
req, err := http.NewRequest("GET", "http://google.com", nil)
req.Header.Set("Accept", "application/json") // Customizando header
```

### 13.4 - Trabalhando com http usando contextos

Utilizando o pacote `context` é possível definir regras de execução para que aquele contexto de código seja interrompido
dependendo da regra definida.
Por exemplo, podemos definir um contexto para que a operação seja cancelada caso uma requisição demorar mais de 1
segundo para ser executada, veja o exemplo abaixo:

```GO
ctx := context.Background() // Criando contexto vazio
ctx, cancel := context.WithTimeout(ctx, time.Second) // Definindo regra de timeout de 1 segundo.

req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com.br", nil) // Requisição http passando o contexto.
```