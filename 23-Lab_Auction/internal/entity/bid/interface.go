package bid_entity

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

type BidRepositoryInterface interface {
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)

	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
}
