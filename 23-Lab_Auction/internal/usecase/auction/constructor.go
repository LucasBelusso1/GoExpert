package auction_usecase

import (
	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	bid_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/bid"
)

type AuctionUseCase struct {
	AauctionRepository auction_entity.AuctionRepositoryInterface
	BidRepository      bid_entity.BidRepositoryInterface
}

func NewAuctionUseCase(
	auctionRepository auction_entity.AuctionRepositoryInterface,
	bidRepository bid_entity.BidRepositoryInterface,
) AuctionUseCaseInterface {
	return &AuctionUseCase{
		AauctionRepository: auctionRepository,
		BidRepository:      bidRepository,
	}
}
