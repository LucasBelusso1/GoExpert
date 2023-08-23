### Many to Many

Para fazer uma relação **n:n**, primeiramente temos que informar nas structs que ambas as structs de produto e categoria
possuem um slice da relação, veja o exemplo:

```GO
type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}
```

Note que precisamos informar ao `gorm` que há uma relação de `many2many` informando a tabela intermediária.

Depois de feita a relação, precisamos criar os registos e consultar, desta forma:

```GO
db.AutoMigrate(&Product{}, &Category{})

Create category
category := Category{Name: "Cozinha"}
db.Create(&category)

category2 := Category{Name: "Eletronicos"}
db.Create(&category2)

db.Create(&Product{
	Name:       "Mouse",
	Price:      1000.00,
	Categories: []Category{category, category2},
})

var categories []Category
err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

if err != nil {
	panic(err)
}

for _, category := range categories {
	fmt.Println(category.Name, ":")
	for _, product := range category.Products {
		println("- ", product.Name, "Serial Number:")
	}
}
```