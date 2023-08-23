### HasOne

A ideia neste caso é fazer um relacionamento **1:1** entre um produto e um serial number, ou seja, cada produto terá
somente um serial number e vice-versa.
Para fazer isso, precisamos declarar uma struct `SerialNumber` que possui o ID de um produto.

```GO
type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}
```

E no produto precisamos informar que ele terá relação com o SerialNumber, desta forma:

```GO
type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}
```

Agora criamos o produto como nos exemplos anteriores e criamos um SerialNumber passando o ID do produto que foi criado:

```GO
db.Create(&SerialNumber{
	Number:    "123456",
	ProductID: 1,
})
```

Ao buscar os dados inseridos, também precisamos fazer o `Preload` da nova tabela que foi criada:

```GO
var products []Product

db.Preload("Category").Preload("SerialNumber").Find(&products)
for _, product := range products {
	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
}
```