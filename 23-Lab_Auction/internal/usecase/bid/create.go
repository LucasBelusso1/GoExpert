package bid_usecase

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	bid_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/bid"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

var bidBatch []bid_entity.Bid

func (bu *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(bu.bidChannel)

		for {
			select {
			case bidEntity, ok := <-bu.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						err := bu.BidRepository.CreateBid(ctx, bidBatch)
						if err != nil {
							logger.Error("error trying to process bid batch list", err)
						}
					}
					return
				}

				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= bu.maxBatchSize {
					err := bu.BidRepository.CreateBid(ctx, bidBatch)
					if err != nil {
						logger.Error("error trying to process bid batch list", err)
					}

					bidBatch = nil
					bu.timer.Reset(bu.batchInsertInterval)
				}
			case <-bu.timer.C:
				err := bu.BidRepository.CreateBid(ctx, bidBatch)
				if err != nil {
					logger.Error("error trying to process bid batch list", err)
				}

				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)
			}

		}
	}()
}

func (bu *BidUseCase) CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError {
	bidEntity, err := bid_entity.NewBidEntity(bidInputDTO.UserId, bidInputDTO.AuctionId, bidInputDTO.Amount)

	if err != nil {
		return err
	}

	bu.bidChannel <- *bidEntity

	return nil
}
