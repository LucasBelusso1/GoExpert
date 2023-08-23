### Realizando consultas com WHERE

Exemplos:

```GO
// Limitando registros
var products []Product
db.Limit(2).Find(&products)

for _, product := range products {
	fmt.Println(product)
}

// Limitando registros e passando um offset
var products []Product
db.Limit(2).Offset(2).Find(&products)

for _, product := range products {
	fmt.Println(product)
}

// WHERE
var products []Product
db.Where("price > ?", 200).Find(&products)

for _, product := range products {
	fmt.Println(product)
}

// WHERE + LIKE
var products []Product
db.Where("name LIKE ?", "%ouse%").Find(&products)

for _, product := range products {
	fmt.Println(product)
}
```