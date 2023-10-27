package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	productUseCase := NewUseCase(db)

	product, err := productUseCase.GetProduct(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
