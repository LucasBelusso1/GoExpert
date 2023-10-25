package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LucasBelusso1/15-SQLC/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")

	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	uuid := uuid.New().String()
	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid,
		Name:        "Backend",
		Description: sql.NullString{String: "Backend description", Valid: true},
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("List of created categories.")
	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          uuid,
		Name:        "Backend updated",
		Description: sql.NullString{String: "Backend updated description", Valid: true},
	})

	fmt.Println("List of updated categories.")
	categories, err = queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	if err != nil {
		panic(err)
	}

	err = queries.DeleteCategory(ctx, uuid)

	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted categories (nothing should be printed).")
	categories, err = queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}
}
