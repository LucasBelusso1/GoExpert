package auction_entity

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

type AuctionRepositoryInterface interface {
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internal_error.InternalError)

	CreateAuction(ctx context.Context, auctionEntity *Auction) *internal_error.InternalError
}
