package auction_controller

import (
	"context"
	"net/http"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/validation"
	auction_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/auction"
	"github.com/gin-gonic/gin"
)

func (ac *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDto auction_usecase.AuctionInputDTO

	err := c.ShouldBindJSON(&auctionInputDto)

	if err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	createAuctionError := ac.AuctionUseCase.CreateAuction(context.Background(), auctionInputDto)

	if createAuctionError != nil {
		restErr := rest_err.ConvertError(createAuctionError)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
