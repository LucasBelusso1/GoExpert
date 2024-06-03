package bid

import (
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/database/auction"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection:        database.Collection("bid"),
		AuctionRepository: auctionRepository,
	}
}
