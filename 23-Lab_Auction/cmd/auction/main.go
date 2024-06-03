package main

import (
	"context"
	"log"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/database/mongodb"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/controller/user_controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("cmd/auction/.env")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx := context.Background()
	_, err = mongodb.NewMongoDbConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	router.GET("/auctions")
	router.POST("/auctions")
	router.POST("/auction/winner/:auctionId")

	router.GET("/bid/:auctionId")
	router.POST("/bid")

	router.GET("/user/:userId")

	router.Run(":8080")
}

func initDependecies() (user user_controller.UserController) {}
