package auction

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	_, err := ar.Collection.InsertOne(ctx, &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	})

	if err != nil {
		logger.Error("Error trying to create auction", err)
		return internal_error.NewInternalServerError("Error trying to create auction")
	}

	return nil
}
