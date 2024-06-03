package main

import (
	"context"
	"log"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/database/mongodb"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/controller/auction_controller"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/controller/bid_controller"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/controller/user_controller"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/database/auction"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/database/bid"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/database/user"
	auction_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/auction"
	bid_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/bid"
	user_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	err := godotenv.Load("cmd/auction/.env")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx := context.Background()
	databaseConn, err := mongodb.NewMongoDbConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	user, auction, bid := initDependecies(databaseConn)

	router.GET("/auctions", auction.FindAuctions)
	router.GET("/auctions/:auctionId", auction.FindAuctionById)
	router.GET("/auction/winner/:auctionId", auction.FindWinningBidByAuctionId)
	router.POST("/auctions", auction.CreateAuction)

	router.GET("/bid/:auctionId", bid.FindBidByAuctionId)
	router.POST("/bid", bid.CreateBid)

	router.GET("/user/:userId", user.FindUserById)

	router.Run(":8080")
}

func initDependecies(database *mongo.Database) (
	userController *user_controller.UserController,
	auctionController *auction_controller.AuctionController,
	bidController *bid_controller.BidController,
) {
	userRepository := user.NewUserRepository(database)
	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
