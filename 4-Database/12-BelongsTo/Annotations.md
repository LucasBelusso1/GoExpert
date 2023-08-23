### BelongsTo

A ideia aqui é fazer um relacionamento entre tabelas, sendo assim primeiramente precisamos criar uma outra struct
para se relacionar com a struct que já existe de products:

```GO
type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
```

Após criar a struct de categoria, precisamos informar na struct de products que cada produto será vinculado a uma
categoria:

```GO
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}
```

Feito o vínculo dentro das structs, agora basta criar as tabelas e criar os registros.
Primeiro devemos criar um registro na tabela de categorias, que existe independente de um produto, e depois precisamos
passar essa categoria para dentro do produto ao criá-lo:

```GO
	db.AutoMigrate(&Product{}, &Category{})

	// Create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Mouse",
		Price:      1000.00,
		CategoryID: 1,
	})
```
Agora para recuperar os dados dos produtos, precisamos fazer um `Preload` das categorias antes de buscarmos os produtos,
do contrário as informações das categorias não serão carregadas.

```GO
	var products []Product

	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
	}
```