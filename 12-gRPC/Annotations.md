### gRPC

Remote Procedure Call (gRPC) é uma ferramenta utilizada para comunições entre sistemas desenvolvida pela Google e
mantida pela CNCF que utiliza HTTP 2 com comunicação bidirecional, ou seja, em uma mesma conexão é possível enviar
múltiplas requests e receber múltiplas responses. É extremamente indicada para a comunicação entre microserviços e
possui suporte oficial em diversas linguagens, como GO, JAVA e C.

### Protocol Buffer (Protobuf)

É um protocolo desenvolvido também pela Google que trafega dados em binário e possui contratos semelhantes ao que existe
com o XML. Nele é possível definir "schemas" de entrada e saída de dados com tipos pré definidos ou tipos criados pelo
próprio desenvolvedor. Este protocolo é mais performático do que JSON, pois é a ideia é que ele trafegue binários, então
ele é consome menos recursos de rede e o processo de serialização e deserialização é mais rápido.

### HTTP/2

O HTTP/2 é uma nova versão do protocolo HTTP que suporta tráfego de binários de forma bidirecional, ou seja, em uma
mesma conexão é possível enviar e receber múltiplas requisições, além de possuir um sistema de push para carregamento
de arquivos de assets por exemplo.

### Formatos de comunicação

`gRPC - API "unary"`: Comunicação via request seguida de response, semelhante ao que existe no HTTP 1
`gRPC - API "Server streaming"`: Neste cenário, enviamos uma requisição para o servidor e este pode mandar múltiplas
respostas para o client processar.
`gRPC - API "Client streaming"`: Semelhante a API server streaming, porém neste caso o client pode enviar múltiplas
requisições para o servidor ir processando.
`gRPC - API "Bi directional streaming"`: Tanto o client quanto o server podem enviar dados continuamente um para o outro
sem necessariamente aguardar por dado.

### REST vs gRPC

`REST`: A comunicação no padrão rest é unidirecional, é em formato text/JSON, possui uma alta latência, não possui
garantia de contrato, não possui suporte a streaming, possui design pré definido (GET, POST, PUT, PATCH, DELETE...)
pensado para um CRUD basicamente.
`gRPC`: Utiliza Protocol buffers, é bidirecional e assíncrono, possui baixa latência, possui contrato (.proto), possui
suporte a streaming, o design é livre, possui geração de código.

### Instalado compilador e plugins

Primeiramente precisamos instalar na máquina o compilador do Protocol Buffer (protoc).

Veja o artigo oficial de [como instalar](https://grpc.io/docs/protoc-installation/).

Em seguida precisamos instalar os plugins para que o protobuf e o grpc funcionem no GO.

Veja o artigo oficial de [como instalar](https://grpc.io/docs/languages/go/quickstart/).

### Fazendo setup do projeto

Para fazer o setup do projeto, copiamos a pasta `/database` do repositório `11-GraphQL` e colamos dentro da pasta
`internal` do nosso projeto atual e executamos o `go mod tidy` para baixar as dependências.

### Crianto protofile

Dica: Instalar extensão [vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)

Dentro da raiz do projeto, criamos a pasta `/proto` e nela criamos o arquivo `course_category.proto` com o seguinte
conteúdo:

```proto
syntax = "proto3"; // Syntaxe utilizada.
package pb; // Indica o nome do pacote (pb = protocol buffers).
option go_package = "internal/pb"; // Cria uma pasta chamada pb dentro de `internal`.

// Message com os campos da categoria.
message Category {
	string id = 1;
	string name = 2;
	string description = 3;
}

// Message com os campos da request para criação de categoria.
message CreateCategoryRequest {
	string name = 1;
	string description = 2;
}

// Message com o conteúdo da response.
message CategoryResponse {
	Category category = 1;
}

// Serviço que possui a função CreateCategory que recebe uma category request e retorna uma category response.
service CategoryService {
	rpc CreateCategory(CreateCategoryRequest) returns (CategoryResponse) {}
}
```

### Fazendo geração de código com protoc

Agora para gerar o código do gRPC e do Protocol Buffers, rodamos o seguinte comando:

```SHELL
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

Este comando criará uma pasta dentro de `internal` chamada `pb` (pré definida no arquivo .proto), e criará 2 arquivos,
um para comunicar-se com o protocol buffers e outro para comunicar-se com o gRPC.

### Implementando CreateCategory

Agora dentro de `/internal` criamos a pasta `/service` com o arquivo `category.go`. Nele adicionaremos o seguinte:

```GO
type CategoryService struct {
	pb.UnimplementedCategoryServiceServer // Retrocompatibilidade com o pb
	CategoryDB database.Category
}

// Construtor
func CreateCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

// Função CreateCategory definida dentro de course_category_grpc.pb.go na parte de "server"
func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)

	if err != nil {
		return nil, err
	}

	categoryPb := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: categoryPb}, nil
}
```

### Criando servidor gRPC

Para criar o servidor, criamos a pasta `/cmd/grpcServer` na raiz do projeto com o arquivo `main.go` com o seguinte
conteúdo:

```GO
func main() {
	db, err := sql.Open("sqlite3", "database.db") // Abre conexão com o banco de dados

	if err != nil {
		panic(err)
	}

	defer db.Close()
	categoryDB := database.NewCategory(db) // Cria uma categoria e passa o banco
	categoryService := service.NewCategoryService(*categoryDB) // Cria um CategoryService que faz parte da interface do server

	grpcServer := grpc.NewServer() // Inicializa o servidor gRPC
	pb.RegisterCategoryServiceServer(grpcServer, categoryService) // Registra o server de cateogoria no servidor do gRPC
	reflection.Register(grpcServer) // Utilizado para o client Evans que será visto mais a frente

	lis, err := net.Listen("tcp", ":50051") // Abre porta TCP para se comunicar com o grpc

	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
```

### Interagindo com Evans

Para começar a interagir com o [Evans](https://github.com/ktr0731/evans), primeiro precisamos instalá-lo (veja como
no link). Em seguida inicializamos nosso servidor na porta 50051 com o comando `go run main.go`.
Com o servidor inicializado podemos chamar o evans da seguinte forma:

```SHELL
evans -r repl
```

Em alguns casos o evans irá puxar tanto o pacote quanto o service dependendo da porta utilizada. Caso isso não ocorra,
é necessário executar os comandos

```SHELL
package pb
```
Para especificar nosso pacote (definido no `course_category.proto`). E em seguida:

```SHELL
service CategoryService
```
Para especificar o serviço que utilizaremos.
Após isso basta fornecer os dados e criar nosso registro.

### Criando categoryList no protofile

Para criar o categoryList, primeiramente editamos o arquivo `.proto` e adicionamos duas novas `message`:

```proto
message blank {}
message CategoryList {
	repeated Category categories = 1;
}
```

No caso do `CategoryList`, informamos antes da propriedade `Category categories` a palavra reservada `repeated`, que
informa que aquele dado irá se repetir indefinidamente.
Além disso, dentro do service `CategoryService`, adicionamos o seguinte:

```proto
service CategoryService {
	rpc CreateCategory(CreateCategoryRequest) returns (Category) {}
	rpc ListCategories(blank) returns (CategoryList) {} // Nova linha
}
```
Aqui entendemos o por que do `message blank{}` mais acima. Adicionamos o `blank` porque no caso do protobuf, não existe
chamada vazia, sendo assim, criamos uma mensagem que não possui propriedades, assim ao executar o `call` não precisamos
informar parâmetros.

Agora executamos o comando para gerar novamente os arquivos do `pb`:

```SHELL
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

Em seguida, precisamos implementar a função `ListCategories` dentro do nosso service, bem como indica no arquivo
`course_category_grpc.pb.go`:

```GO
func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesPB []*pb.Category
	for _, category := range categories {
		categoriesPB = append(categoriesPB, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoriesPB}, nil
}
```

Feito isso já é possível chamar o ListCategories via `evans`.


### Buscando uma categoria

Aqui basicamente segue o mesmo processo do tópico **Criando categoryList no protofile**.
Arquivo proto:

```proto
message CategoryGetRequest {
	string id = 1;
}

service CategoryService {
	rpc CreateCategory(CreateCategoryRequest) returns (Category) {}
	rpc ListCategories(blank) returns (CategoryList) {}
	rpc getCategory(CategoryGetRequest) returns (Category) {} // Nova linha
}
```

Arquivo service do go:

```GO
func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}
```

### Trabalhando com streams

Agora o processo muda um pouco, veja o arquivo `.proto`:

```proto
service CategoryService {
	rpc CreateCategory(CreateCategoryRequest) returns (Category) {}
	rpc CreateCategoryStream(stream CreateCategoryRequest) returns (CategoryList) {} // Nova linha
	rpc ListCategories(blank) returns (CategoryList) {}
	rpc getCategory(CategoryGetRequest) returns (Category) {}
}
```

Criamos então o `rpc` **CreateCategoryStream** informando que a request será uma stream de dados e o response vai ser
uma lista de categorias. Isso significa que a conexão ficará aberta recebendo indefinidas categorias e quando ela chegar
ao fim, a lista de categorias será retornada.

Veja como fica o código dentro do service:

```GO
func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}
```

Perceba que agora recebemos como parâmetro uma interface do tipo `pb.CategoryService_CreateCategoryStreamServer` que
possui duas funções:

- `Recv()`: Que recebe os dados que estão sendo inseridos.
- `SendAndClose()`: Que retorna os dados criados e fecha a conexão.

Então criamos um looping infinito para receber todos os dados do stream e validamos, caso a conexão seja interrompida
(`io.EOF`) retornamos todos os dados.

### Trabalhando com streams bidirecionais

Aqui é semelhante ao que foi feito no passo anterior.
Veja o arquivo proto:

```proto
service CategoryService {
	rpc CreateCategory(CreateCategoryRequest) returns (Category) {}
	rpc CreateCategoryStream(stream CreateCategoryRequest) returns (CategoryList) {}
	rpc CreateCategoryStreamBidirectional(stream CreateCategoryRequest) returns (stream Category) {} // Nova linha
	rpc ListCategories(blank) returns (CategoryList) {}
	rpc getCategory(CategoryGetRequest) returns (Category) {}
}
```

Veja o arquivo GO:

```GO
func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)

		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})

		if err != nil {
			return err
		}
	}
}
```

Aqui a diferença é que ao invés de enviarmos todos os dados processador de uma vez ao fechar a conexão, fazemos o envio
do dado no momento em que ele é informado e processado.