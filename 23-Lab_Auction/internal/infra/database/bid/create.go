package bid

import (
	"context"
	"sync"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/logger"
	auction_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/auction"
	bid_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/bid"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()
			// TODO: Validar quando o canal vai fechar caso hajam muitos registros do mesmo banco
			// (talvez usando um canal compartilhado entre a operação de updateAuction e CreateBid com uma flag que identifique o status)
			auctionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)

			if err != nil {
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			_, error := bd.Collection.InsertOne(ctx, &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			})

			if error != nil {
				logger.Error("Error trying to insert bid", error)
				return
			}
		}(bid)
	}

	wg.Wait()
	return nil
}
