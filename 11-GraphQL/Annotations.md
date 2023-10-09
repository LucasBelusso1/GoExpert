### Gerando esqueleto do servidor GraphQL

Para criar o projeto utilizando GraphQL, utilizaremos o [gqlgen](https://gqlgen.com/). Para isso, primeiro rodamos o
comando:

```SHELL
printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools.go
```

Isso criará um arquivo `tools.go` com alguns pacotes. Para importar as dependências deste pacote executademos o comando
`go mod tidy`. Em seguida executamos um `init` do pacote `github.com/99designs/gqlgen` que foi anteriormente baixado:

```SHELL
go run github.com/99designs/gqlgen init
```

Isso fará com que algumas pastas e arquivos sejam criados para começarmos o projeto GraphQL.
Para executar o servidor GraphQL podemos rodar o comando `go run server.go`, que nos dará acesso ao playground do
GraphQL em `https://localhost:8080`.

### Criando schema GraphQL

Aqui vamos criar um schema que vai possuir os models, os inputs, a query e a mutation, desta forma:

```GRAPHQLS
// Model
type Category {
  id: ID!
  name: String!
  description: String
  courses: [Course!]!
}

// Model
type Course {
  id: ID!
  name: String!
  description: String
  category: Category!
}

// Input
input NewCategory {
  name: String!
  description: String
}

// Input
input NewCourse {
  name: String!
  description: String
  categoryId: ID!
}

// Query
type Query {
  categories: [Category!]!
  courses: [Course!]!
}

// Mutation
type Mutation {
  createCategory(input: NewCategory!): Category!
  createCourse(input: NewCourse!): Course!
}
```

Agora executamos o comando:

```SHELL
go run github.com/99designs/gqlgen generate
```

Que gerará novamente os arquivos porém agora considerando nosso schema de cursos e categorias.

### Criando resolver para Category

Para criar o resolver de Category, primeiramente separamos o server dentro de `cmd/server`, depois, criamos a pasta
`internal/database` e nele criamos o arquivo `category.go` que vai conter o seguinte código:

```GO
type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
```

Neste caso criamos a struct de categoria, criamos a função `NewCategory()` para atuar como construtor e receber o banco
de dados como parâmetro, e criamos a função `Create()`.
Agora dentro do arquivo `resolver.go` gerado pelo pacote que importamos do `glqgen`, injetamos nossa struct desta forma:

```GO
type Resolver struct{
	CategoryDB *database.Category
}
```

Agora dentro do `schema.resolvers.go`, na função `Create()` inserimos o seguinte:

```GO
// Criamos nossa categoria com as informações do input
category, err := r.CategoryDB.Create(input.Name, *input.Description)

// Validamos erro
if err != nil {
  return nil, err
}

// Criamos a categoria do model gerado pelo gqlgen e retornamos.
return &model.Category{
  ID:          category.ID,
  Name:        category.Name,
  Description: &category.Description,
}, nil
```

### Persistindo categoria via playground

Agora para podermos inserir no banco uma categoria através do GraphQL, precisamos primeiro adicionar uma conexão com um
banco de dados, que neste caso será o `sqlite3`. Dentro de `server.go` adicionamos o seguinte código:

```GO
db, err := sql.Open("sqlite3", "./data.db")

if err != nil {
  log.Fatalf("Failed to open database: %v", err)
}
defer db.Close()
```

Em seguida, ainda dentro de `server.go` criamos a nossa categoryDB e injetamos ela nos resolvers:

```GO
categoryDB := database.NewCategory(db)
srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
  CategoryDB: categoryDB,
}}))
```

Agora basta criarmos a nossa tabela dentro do sqlite e rodar nosso servidor. Para criar o registro através do GraphQL
fazemos da seguinte forma:

```
mutation createCategory {
	createCategory(input: {name: "Tecnologia", description: "Cursos de tecnologia"}) {
    id
    name
    description
  }
}
```

### Fazendo query de categorias

Para buscar as categorias seguimos um processo parecido. Criamos a função `FindAll` dentro do resolver de categoria,
desta forma:

```GO
func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		category := Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.Description)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
```

Agora editamos o schema resolver para buscar as categorias pelo resolver:

```GO
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var categoriesModel []*model.Category
	for _, category := range categories {
		categoriesModel = append(categoriesModel, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}

	return categoriesModel, nil
}
```

### Criando resolvers de courses

Para os courses basicamente seguimos os mesmos passos das categorias.

### Dados encadeados.

Para encadear os dados, primeramente, dentro de `model`, criamos dois novos arquivos, `category.go` e `course.go` para
guardar nossos models que estão presentes dentro de `models_gen.go`. Em seguida editamos o model de **Category** e
removemos o array de courses e no arquivo `gqlgen.yml`, nos models, adicionamos o seguinte:

```YML
  Category:
    model:
      - github.com/LucasBelusso1/11-GraphQL/graph/model.Category
  Course:
    model:
      - github.com/LucasBelusso1/11-GraphQL/graph/model.Course
```

E rodamos novamente o comando `go run github.com/99designs/gqlgen generate`. Isso fará com que uma nova função seja
criada dentro do `schema.resolvers.go`, nesta função vamos criar a lógica para buscar um curso pelo ID da categoria.

Agora dentro de `database/course.go` criamos a função `FindByCategoryID()`:

```GO
func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description, category_id FROM courses where category_id = $1", categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		course := Course{}
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}
```
Depois, dentro do `schema.resolvers.go`, implementamos a busca utilizando a função que criamos:

```GO
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryID(obj.ID)

	if err != nil {
		return nil, err
	}

	var coursesModel []*model.Course
	for _, course := range courses {
		coursesModel = append(coursesModel, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
		})
	}

	return coursesModel, nil
}
```

Desta forma, será possível buscar todos os cursos e suas categorias com o seguinte comando no GraphQL

```GRAPHQL
query queryCategoriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
    }
  }
}
```

### Finalizando encadeamento de categorias.

Agora, vamos buscar os cursos e suas respectivas categorias. Para isso, vamos executar um processo semelhante ao
tópico anterior.

Vamos até os models e removemos o campo de `Category` do model de cursos, e em seguida rodamos novamente o comando
`go run github.com/99designs/gqlgen generate`. Novamente isso gerará uma struct nova no `schema.resolvers.go` com a
função `Category`.
Agora precisaremos ir até o resolver de cateogria e criar a função para buscar a categoria pelo curso:

```GO
func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var id, name, description string

	query := "SELECT c.id, c.name, c.description FROM categories c JOIN courses co ON c.id = co.category_id where co.id = $1"
	err := c.db.QueryRow(query, courseID).Scan(&id, &name, &description)

	if err != nil {
		return Category{}, nil
	}

	return Category{ID: id, Name: name, Description: description}, nil
}
```

Em seguida, no schema resolver implementamos a busca de fato:

```GO
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	course, err := r.CategoryDB.FindByCourseID(obj.ID)

	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          course.ID,
		Name:        course.Name,
		Description: &course.Description,
	}, nil
}
```

Agora é possível obter os cursos e suas respectivas categorias, desta forma:

```GRAPHQL
query queryCourseWithCategory {
  courses {
    id
    name
    description
    category {
      id
      name
      description
    }
  }
}
```