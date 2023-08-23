### Configurando e criando operações

Para se conectar ao banco de dados utilizando o GORM, primeiro temos que importá-lo ao projeto utilizando o comando
`go get -u gorm.io/gorm` e depois precisamos importar o driver do banco que iremos utilizar, no nosso caso estaremos
utilizando mysql, então o import é assim: `go get -u gorm.io/driver/mysql`.

Feito a importação dos módulos, agora precisamos conectar com o banco. Para isso precisamos chamar a função `Open` do
pacote `gorm` passando o driver e suas configurações, desta maneira:

```GO
dsn := "root:root@tcp(localhost:3306)/goexpert"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```

Após realizar a conexão com o banco, é possível definir structs com anotações (semelhante ao que existe para JSON)
para criar as tabelas do banco:

```GO
type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

db.AutoMigrate(&Product{})
```

Feito a criação da tabela, já podemos começar a inserir registros, individualmente ou em lote, veja o exemplo:

```GO
// Create
db.Create(&Product{
	Name:  "Notebook",
	Price: 1000.0,
})

//Create in batch
products := []Product{
	{Name: "Notebook", Price: 1000.0},
	{Name: "Mouse", Price: 100.0},
	{Name: "Keyboard", Price: 200.0},
}

db.Create(&products)
```