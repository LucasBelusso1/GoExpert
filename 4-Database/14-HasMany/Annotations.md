### HasMany

Neste caso temos a relação de **1:n**, ou seja, o produto possui apenas uma categoria, enquanto que a categoria pode ter
**n** produtos, veja o exemplo:

```GO
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

db.Create(&category)

db.Create(&Product{
	Name:       "Mouse",
	Price:      1000.00,
	CategoryID: 2,
})

var categories []Category
err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

if err != nil {
	panic(err)
}

for _, category := range categories {
	fmt.Println(category.Name, ":")
	for _, product := range category.Products {
		println("- ", product.Name)
	}
}
```

Neste caso, na struct declaramos que a categoria terá um slice de produtos, e ao consultarmos a categoria, carregamos
a tabela de produtos e aplicamos um `Find`.