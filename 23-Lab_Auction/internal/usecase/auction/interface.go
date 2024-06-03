package auction_usecase

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

type AuctionUseCaseInterface interface {
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)

	CreateAuction(ctx context.Context, auctionInputDto AuctionInputDTO) *internal_error.InternalError
}
