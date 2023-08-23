package main

import (
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
}
