package auction_usecase

import (
	"context"

	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInputDto AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.NewAuction(
		auctionInputDto.ProductName,
		auctionInputDto.Category,
		auctionInputDto.Description,
		auction_entity.ProductCondition(auctionInputDto.Condition),
	)

	if err != nil {
		return err
	}

	err = au.AauctionRepository.CreateAuction(ctx, auction)

	if err != nil {
		return err
	}

	return nil
}
