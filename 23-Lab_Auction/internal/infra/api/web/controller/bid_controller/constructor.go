package bid_controller

import bid_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/bid"

type BidController struct {
	BidUseCase bid_usecase.BidUseCase
}

func NewBidController(bidUseCase bid_usecase.BidUseCase) *BidController {
	return &BidController{BidUseCase: bidUseCase}
}
