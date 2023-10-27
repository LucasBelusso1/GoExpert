### DI (Dependency Injection)

Qual o problema?

Em um projeto de médio/grande porte, é possível notar diversas dependencias em diversos níveis da aplicação. Por exemplo,
em alguns projetos anteriores, utilizamos o padrão `UseCase -> Repository -> Entity` no qual primeiro precisávamos criar
o `db`, em seguida criar um `Repository` passando o `db` (logo `db` é uma dependência de `repository`). Em seguida
criávamos o `UseCase` passando o `Repository` (logo `Repository` é uma dependência de `UseCase`). E isso pode escalar
em diversos níveis, dificultando por exemplo os testes da aplicação.

### Google Wire

Primeiramente é preciso instalar o Wire, veja como acessando o [repositório do projeto](https://github.com/google/wire).

Agora criaremos o arquivo `wire.go` com o seguite código:

```GO
//go:build wireinject

// + build wireinject

package main

import (
	"database/sql"

	"github.com/LucasBelusso1/17-DI/product"
	"github.com/google/wire"
)

func NewUseCase(db *sql.DB) *product.ProductUseCase { // Função que retornará nosso objeto
	wire.Build(
		// Lista de dependências necessárias
		product.NewProductRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
```

No código acima, declaramos que ao buildar o projeto, este arquivo não será utilizado, somente será utilizado pelo CLI
do wire.

Agora na nossa main podemos chamar da seguinte forma:

```GO
func main() {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	productUseCase := NewUseCase(db) // Cria o UseCase com todas as dependências necessárias
	product, err := productUseCase.GetProduct(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
```

### Trabalhando com Set's de dependências

Em um cenário em que queremos utilizar as mesmas dependências em lugares diferentes, no arquivo `wire.go`, podemos
adicionar o seguinte:

```GO
var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	// ... Lista de dependências
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency, // Utilizando a lista
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}
```

Quando queremos trabalhar com interfaces e precisamos que essa interface sempre assuma um valor específico, podemos
fazer o seguinte:

```GO
wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
```

Assim informamos ao `wire` que sempre que um `product.ProductRepositoryInterface` for uma dependência de algo, ele
utilizará `*product.ProductRepository` no lugar.

**IMPORTANTE:** Para executar o script utilizando o wire, é necessário rodar `go run main.go wire_gen.go`, pois ambos
pertencem ao pacote `main`.