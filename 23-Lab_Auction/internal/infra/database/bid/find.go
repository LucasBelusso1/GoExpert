package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	bid_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/bid"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)

	if err != nil {
		errorMessage := fmt.Sprintf("Error trying to find bids by auctionId, %s", auctionId)
		logger.Error(errorMessage, err)
		return nil, internal_error.NewNotFoundError(errorMessage)
	}

	var bidEntitiesMongo []BidEntityMongo

	err = cursor.All(ctx, &bidEntitiesMongo)

	if err != nil {
		errorMessage := fmt.Sprintf("Error trying to find bids by auctionId, %s", auctionId)
		logger.Error(errorMessage, err)
		return nil, internal_error.NewInternalServerError(errorMessage)
	}

	var bidEntities []bid_entity.Bid

	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bidEntityMongo BidEntityMongo
	err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo)

	if err != nil {
		logger.Error("Error trying to find the auction winner", err)
		return nil, internal_error.NewInternalServerError("Error trying to find the auction winner")
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
