package bid_controller

import (
	"context"
	"net/http"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	"github.com/LucasBelusso1/23-Lab_Auction/internal/infra/api/web/validation"
	bid_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/bid"
	"github.com/gin-gonic/gin"
)

func (bc *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO

	err := c.ShouldBindJSON(&bidInputDTO)

	if err != nil {
		restErr := validation.ValidateErr(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	createBidError := bc.BidUseCase.CreateBid(context.Background(), bidInputDTO)

	if createBidError != nil {
		restErr := rest_err.ConvertError(createBidError)
		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)
}
