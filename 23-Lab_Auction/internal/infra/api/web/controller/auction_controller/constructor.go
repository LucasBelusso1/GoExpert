package auction_controller

import auction_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/auction"

type AuctionController struct {
	AuctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{AuctionUseCase: auctionUseCase}
}
