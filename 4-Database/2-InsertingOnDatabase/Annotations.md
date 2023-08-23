### Inserindo dados no banco

Para manipular os dados no banco primeiramente temos que abrir uma conexão com ele, utilizando o pacote `database/sql`.
Na função main, iniciamos a conexão desta forma:

```GO
sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
```

Para utilizar o mysql ou qualquer outro banco com o GO, precisamos importar o pacote de drivers manualmente, desta
forma:

```GO
import _ "github.com/go-sql-driver/mysql"
```

No caso acima o underline, ou "Blank identifier" precisa ser declarado pois ele indica para o GO que não utilizaremos
diretamente o pacote, mas precisamos que ele seja carregado.

Agora precisamos preparar a nossa query para evitar vulnerabilidades com **SQL INJECTION**, desta forma:

```GO
stmt, err := db.Prepare("INSERT INTO products(id, name, price) VALUES (?, ?, ?)")
```

Feito isso agora precisamos criar uma `struct` e passar seus valores para a função `Exec()` da variável `stmt` que foi
criada acima:

```GO
product := newProduct("Notebook", 1899.90)
_, err = stmt.Exec(product.ID, product.Name, product.Price)
```

Desta forma, os valores indicados pela interrogação na preparação da query serão substituídos pelos valores passados
para a função `Exec()` **em ordem**.