package bid_controller

import bid_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/bid"

type BidController struct {
	BidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{BidUseCase: bidUseCase}
}
