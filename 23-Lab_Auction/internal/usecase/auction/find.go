package auction_usecase

import (
	"context"

	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	bid_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/bid"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := au.AauctionRepository.FindAuctionById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {
	auctionsEntities, err := au.AauctionRepository.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)

	if err != nil {
		return nil, err
	}

	var auctionsOutputDtos []AuctionOutputDTO
	for _, auctionEntity := range auctionsEntities {
		auctionsOutputDtos = append(auctionsOutputDtos, AuctionOutputDTO{
			Id:          auctionEntity.Id,
			ProductName: auctionEntity.ProductName,
			Category:    auctionEntity.Category,
			Description: auctionEntity.Description,
			Condition:   ProductCondition(auctionEntity.Condition),
			Status:      AuctionStatus(auctionEntity.Status),
			Timestamp:   auctionEntity.Timestamp,
		})
	}

	return auctionsOutputDtos, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError) {
	auction, err := au.AauctionRepository.FindAuctionById(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	bidWinning, err := au.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	auctionOutputDTO := AuctionOutputDTO{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	if bidWinning == nil {
		return &WinningInfoOutputDTO{
			Auction: auctionOutputDTO,
			Bid:     nil,
		}, nil
	}

	bidOutputDTO := &bid_usecase.BidOutputDTO{
		Id:        bidWinning.Id,
		UserId:    bidWinning.UserId,
		AuctionId: bidWinning.AuctionId,
		Amount:    bidWinning.Amount,
		Timestamp: bidWinning.Timestamp,
	}

	return &WinningInfoOutputDTO{
		Auction: auctionOutputDTO,
		Bid:     bidOutputDTO,
	}, nil
}
