package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	bid_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/bid"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

type BidUseCase struct {
	BidRepository bid_entity.BidRepositoryInterface

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

type BidUseCaseInterface interface {
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)

	CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError
}

func NewBidUseCase(bidRepository bid_entity.BidRepositoryInterface) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	bidUseCase := &BidUseCase{
		BidRepository:       bidRepository,
		timer:               time.NewTimer(maxSizeInterval),
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: maxSizeInterval,
		bidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)

	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	batchSize, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))

	if err != nil {
		return 5
	}

	return batchSize
}
