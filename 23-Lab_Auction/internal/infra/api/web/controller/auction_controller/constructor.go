package auction_controller

import auction_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/auction"

type AuctionController struct {
	AuctionUseCase auction_usecase.AuctionUseCase
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCase) *AuctionController {
	return &AuctionController{AuctionUseCase: auctionUseCase}
}
