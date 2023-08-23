package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.0,
	// })

	//Create in batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000.0},
	// 	{Name: "Mouse", Price: 100.0},
	// 	{Name: "Keyboard", Price: 200.0},
	// }

	// db.Create(&products)

	//Select one
	// var product Product

	// db.First(&product, 1) // Selecionando o primeiro registro retornado
	// fmt.Println(product)

	// // Select with where
	// db.First(&product, "name = ?", "Mouse") // Selecionando o produto cujo nome seja igual a "Mouse"
	// fmt.Println(product)

	// Select all
	// var products []Product
	// db.Find(&products)

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(2).Find(&products) // Limitando registros

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products) // Limitando registros e passando um offset

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// // WHERE
	// var products []Product
	// db.Where("price > ?", 200).Find(&products) // Cláusula WHERE

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// // WHERE + LIKE
	// var products []Product
	// db.Where("name LIKE ?", "%ouse%").Find(&products) // Cláusula WHER + LIKE

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(p)

	var p2 Product
	db.First(&p2, 1)

	fmt.Println(p2.Name)

	db.Delete(p2)
}
