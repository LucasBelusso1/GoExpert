### Migrations

Pacote `https://github.com/golang-migrate/migrate`.
Na documentação da CLI do pacote acima, vamos encontrar o passo a passo de instalação em todos as plataformas. A partir
dai basta seguir a que for referente ao SO.

Feito a instalação rodamos o seguinte comando:

```SHELL
migrate create -ext sql -dir sql/migrations -seq init
```

No comando acima, criamos uma migration informando a extensão (`-ext`) sendo "sql", no diretório (`-dir`)
`sql/migrations` que utilizará uma sequência (`-seq`) com o nome `init`.

Este comando criará a pasta `sql/migrations` com dois arquivos:

`000001_init.up.sql`: Que vai conter a criação do nosso banco de dados.
`000001_init.down.sql`: Que será responsável por apagar tudo o que fizermos no arquivoup.

Agora no arquivo `000001_init.up.sql` vamos inserir o seguinte:

```SQL
CREATE TABLE IF NOT EXISTS categories (
	id VARCHAR(36) NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT
);

CREATE TABLE IF NOT EXISTS courses (
	id VARCHAR(36) NOT NULL PRIMARY KEY,
	category_id VARCHAR(36) NOT NULL,
	name TEXT NOT NULL,
	description TEXT,
	price DECIMAL(10,2) NOT NULL,
	FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

E no arquivo `000001_init.down.sql` adicionamos o seguinte:

```SQL
DROP TABLE IF EXISTS courses;
DROP TABLE IF EXISTS categories;
```

Criamos então um arquivo `docker-compose.yaml` com um container mysql:

```YAML
version: '3'
services:
  mysql:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: courses
      MYSQL_PASSWORD: root
    ports:
      - "3306:3306"
```

E subimos ele rodando `docker compose up -d`.

Em seguida rodamos o seguinte comando:

```SHELL
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
```

Isso fará com que o que está no arquivo `000001_init.up.sql` seja executado.

Em seguida executamos o mesmo comando, porém trocando o `up` por `down`.

```SHELL
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down
```

### SQLX

Pacote `https://github.com/jmoiron/sqlx`.
Trata-se de um pacote que não faz a função de um ORM, mas que ajuda na hora de executar queries dentro do GO, auxiliando
para fazer scans, sanitizar, entre outras funções.

### SQLC

Pacote `https://github.com/sqlc-dev/sqlc`
Primeiramente precisamos fazer a [instalação do SQLC](https://docs.sqlc.dev/en/stable/overview/install.html).

Feita a instalação, agora vamos criar o arquivo de configuração do SQLC, criando o arquivo `sqlc.yaml`:

```YAML
version: "2" # Versão
sql:
- schema: "sql/migrations" # Aonde estarão nossos schemas
  queries: "sql/queries" # Aonde ficarão armazenadas nossas queries
  engine: "mysql" # Qual banco utilizaremos (apenas suporta mysql, postgres e sqlite (em desenvolvimento))
  gen: # Em qual linguagem irá gerar
    go: # Linguagem GO
      package: "db" # Qual pacote do nosso projeto será usado
      out: "internal/db" # Aonde estará este pacote
```

Agora vamos criar, dentro de `/sql`, a pasta `/queries` com o arquivo `query.sql` com o seguinte conteúdo:

```SQL
-- name: ListCategories :many
SELECT * FROM categories;
```

No código acima, usamos algo parecido com uma `annotation`, para informar que o nome da nossa query é `ListCategories` e
que ela retorna muitos resultados.
Agora rodamos o seguinte comando:

```SHELL
sqlc generate
```

Isso fará com que 3 arquivos sejam gerados:

`db.go`: Que terá a interface para manipulação dos dados.
`models.go`: Que terá a estrutura de nossas tabelas.
`query.sql.go`: Que terá o código pronto para ser executado de acordo com a query que criamos em `query.db`.

### Criando CRUD

Agora vamos adicionar mais SQL's ao nosso `query.db` para termos todas as operações de um CRUD:

```SQL
-- name: GetCategory :one
SELECT * FROM categories
WHERE id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description)
VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories
SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;
```

E novamente vamos rodar o comando `sqlc generate`.

Agora podemos trabalhar com o nosso CRUD.
Criamos então a pasta `/cmd/runSQLC` e criamos o arquivo `main.go` com o seguinte código:

```GO
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucasBelusso1/15-SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	uuid := uuid.New().String()
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("List of created categories.")
	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          uuid,
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend updated description", Valid: true},
	})

	fmt.Println("List of updated categories.")
	categories, err = queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, uuid)

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted categories (nothing should be printed).")
	categories, err = queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
```

### Transações

Aqui primeiro precisamos definir a query de inserção de cursos, no `query.db`, desta forma:

```SQL
INSERT INTO courses (id, name, description, category_id)
VALUES(?,?,?,?);
```

Depois criamos uma nova `main.go` dentro de `/runSQLCTX`, copiamos o código da `main.go` criada anteriormente e apagamos
tudo com exceção da criação da conexão e do contexto.

Agora inserimos o seguinte código:

```GO
type CourseDB struct { // Struct que recebe a conexão com o banco e as queries
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB { // Construtor
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil) // Começa a transação
	if err != nil {
		return err
	}

	q := db.New(tx) // Cria um novo objeto Queries passando a transação
	err = fn(q) // Executa a função fn() recebida por parâmetro que possui as Queries a serem executadas.

	if err != nil {
		errRb := tx.Rollback() // Em caso de erro, faz o rollback das Queries
		if errRb != nil {
			return fmt.Errorf("error on roolback: %v, original error: $w", errRb, err) // Concatena ambos os erros de rollback e o erro da Query
		}
		return err
	}

	err = tx.Commit() // Commita as alterações das Queries executadas pela funcão fn().

	if err != nil {
		return err
	}

	return nil
}
```

Agora criamos a função que vai executar as queries dentro da transaction:

```GO
func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error { // Chama a função criada anteriormente e passa uma funcão anônima
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{ // Criação da categoria
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{ // Criação da do curso
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
```

Agora na `main.go`, inserimos o seguinte código:"

```GO
ctx := context.Background()
dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

if err != nil {
	panic(err)
}
defer dbConn.Close()

courseArgs := CourseParams{
	ID:          uuid.NewString(),
	Name:        "GO",
	Description: sql.NullString{String: "Golang course"},
}
categoryArgs := CategoryParams{
	ID:          uuid.NewString(),
	Name:        "Backend",
	Description: sql.NullString{String: "Backend courses"},
}

courseDb := NewCourseDB(dbConn)
err = courseDb.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)

if err != nil {
	panic(err)
}
```

Ao executar este arquivo, um erro será retornado na criação do curso, pois se repararmos bem, ao editarmos o arquivo
`query.sql`, no `INSERT` faltou informar o campo te `price` que é um `NOT NULL`, sendo assim, ao tentar inserir o dado
no banco um erro é retornado pois precisamos informar o campo de `price`.
Neste caso, como todas as nossas queries estavam dentro de uma `transaction`, por conta do erro retornado ao tentar
criar um `course`, toda a transação sofreu `rollback`, e por este motivo a categoria também não é criada.

### Ajustando queries

Para ajustar o campo de preço, primeiro precisamos acessar o `sqlc.yaml` e atualizar para o seguinte:

```YAML
version: "2"
sql:
- schema: "sql/migrations"
  queries: "sql/queries"
  engine: "mysql"
  gen:
    go:
      package: "db"
      out: "internal/db"
      overrides: # Transforma dados
        - db_type: "decimal" # de decimal no banco
          go_type: "float64" # para float64 no GO.
```

Depois rodamos `sqlc generate` novamente. Isso fará com que no model criado pelo SQLC, o campo de `price` passe de
`string` para `float64`. Agora é só fornecer os dados nas structs e rodar novamente o script.

### Exibindo dados com JOIN

Agora para fazer uma consulta com JOIN, basta editarmos o `query.sql` e seguir os mesmos passos anteriores:

```SQL
-- name: ListCourses :many
SELECT c.*, ca.name AS category_name
FROM courses AS c
JOIN categories AS ca ON c.category_id = ca.id;
```
Na `main.go`:

```GO
ctx := context.Background()
dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

if err != nil {
	panic(err)
}
defer dbConn.Close()

queries := db.New(dbConn)

courses, err := queries.ListCourses(ctx)

if err != nil {
	panic(err)
}

for _, course := range courses {
	fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f\n",
		course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
}
```