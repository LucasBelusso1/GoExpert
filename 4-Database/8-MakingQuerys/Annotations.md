### Realizando as primeiras consultas

Utilizando um ORM fica extremamente simples de realizar consultas, veja os exemplos:

```GO
//Select one
var product Product

db.First(&product, 1) // Selecionando o primeiro registro retornado
fmt.Println(product)

// Select with where
db.First(&product, "name = ?", "Mouse") // Selecionando o produto cujo nome seja igual a "Mouse"
fmt.Println(product)

// Select all
var products []Product
db.Find(&products) // Selecionando todos os produtos da tabela

for _, product := range products {
	fmt.Println(product)
}
```