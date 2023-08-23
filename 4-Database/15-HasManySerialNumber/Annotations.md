### Pegadinha no has many

Neste caso a ideia é vincular um `SerialNumber` a um produto e buscar este SerialNumber através de uma consulta em cima
das categorias. Para isso, primeiramente precisamos repor a struct do SerialNumber e relacioná-la com `Product`:

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

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}
```

Agora vamos tentar buscar o SerialNumber pela categoria:

```GO
var categories []Category
err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

if err != nil {
	panic(err)
}

for _, category := range categories {
	fmt.Println(category.Name, ":")
	for _, product := range category.Products {
		println("- ", product.Name, "Serial Number:", product.SerialNumber.Number)
	}
}
```

Neste caso haverá um erro informando que não é possível relacionar categorias com SerialNumber pois não existe essa
relação. Sendo assim, para resolver o problema, precisamos relacionar categorias e SerialNumber através do produto,
desta forma:

```GO
err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
```

Ou então podemos simplificar carregando somente `Products.SerialNumber`:

```GO
err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
```