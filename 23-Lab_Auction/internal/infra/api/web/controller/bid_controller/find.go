package bid_controller

import (
	"context"
	"net/http"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (bc *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	err := uuid.Validate(auctionId)

	if err != nil {
		errRest := rest_err.NewBadRequestError("InvalidFields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auctionWinningData, findAuctionError := bc.BidUseCase.FindBidByAuctionId(context.Background(), auctionId)

	if findAuctionError != nil {
		errRest := rest_err.ConvertError(findAuctionError)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionWinningData)
}
