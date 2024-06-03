package auction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo

	err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("Auction not found with this ID: %s", id), err)
			return nil, internal_error.NewNotFoundError(fmt.Sprintf("Auction not found with this ID: %s", id))
		}

		logger.Error("Error trying to find user by id", err)
		return nil, internal_error.NewNotFoundError("Error trying to find user by id")
	}

	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category, productName string,
) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}
	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewNotFoundError("Error trying to find auctions")
	}

	defer cursor.Close(ctx)

	var auctionsEntityMongo []AuctionEntityMongo
	err = cursor.All(ctx, &auctionsEntityMongo)

	if err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewNotFoundError("Error trying to find auctions")
	}

	var auctionsEntity []auction_entity.Auction

	for _, auctionMongo := range auctionsEntityMongo {
		auctionsEntity = append(auctionsEntity, auction_entity.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Status:      auctionMongo.Status,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctionsEntity, nil
}
