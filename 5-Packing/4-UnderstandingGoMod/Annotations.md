### Entendendo go mod

Podemos utilizar dependências externas de duas formas:

1. Utilizando `go mod tidy` -> Neste caso basta utilizarmos algum pacote em nosso código e rodar o comando `go mod tidy`
para que ele faça a importação.

```GO
import ("github.com/google/uuid")
```

2. Também podemos importar o pacote através do `go get`, que fará a mesma função do `go mod tidy` acima.

```SHELL
go get github.com/google/uuid
```

Ao utilizar alguma dependência externa, é criado o arquivo `go.sum`, que é responsável por fazer o **lock** da versão
dos pacotes que serão importados e suas dependências.