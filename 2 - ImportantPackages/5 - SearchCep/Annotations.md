### Busca CEP

Neste projeto utilizamos muitas das coisas que vimos para gerar um script que busca o CEP por meio de um valor fornecido
via parâmetro no terminal.
Para fazer a chamada HTTP foi utilizado o pacote `net/http` com a função `Get()`.
Para buscar o parâmetro do terminal foi utilizado o pacote `os` pegando a propriedade `Args`.
Utilizou-se o pacote `io` com a função `ReadAll()` para ler o response body.
Utilizou-se o pacote `json` para encodar e decodar o JSON retornado pela API.
