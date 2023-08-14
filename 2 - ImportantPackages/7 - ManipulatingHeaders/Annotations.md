### Manipulando Headers

É possível retornar diferentes respostas HTTP dependendo da necessidade, utilizando o `ResponseWriter` com a função
`WriteHeader()` passando alguma das propriedades disponíveis no pacote `net/http`.
Assim como também é possível definir headers para serem retornados utilizando `.Header().Set()` ou então pegar
parâmetros enviados na URL utilizando o `URL.Query().Get()` do `Request`.