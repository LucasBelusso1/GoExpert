package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LucasBelusso1/11-GraphQL/graph"
	"github.com/LucasBelusso1/11-GraphQL/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
