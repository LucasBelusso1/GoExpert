package main

import (
	"database/sql"
	"net"

	"gihub.com/LucasBelusso1/12-GRPC/internal/database"
	"gihub.com/LucasBelusso1/12-GRPC/internal/pb"
	"gihub.com/LucasBelusso1/12-GRPC/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051") // Abre porta TCP para se comunicar com o grpc

	if err != nil {
		panic(err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
