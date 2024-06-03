package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	auction_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/auction"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (ac *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status ")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, err := strconv.Atoi(status)

	if err != nil {
		errRest := rest_err.NewBadRequestError("Invalid status")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionsData, findAuctionsErr := ac.AuctionUseCase.FindAuctions(
		context.Background(),
		auction_usecase.AuctionStatus(statusNumber),
		category,
		productName,
	)

	if findAuctionsErr != nil {
		errRest := rest_err.ConvertError(findAuctionsErr)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionsData)
}

func (ac *AuctionController) FindAuctionById(c *gin.Context) {
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

	auctionData, findAuctionError := ac.AuctionUseCase.FindAuctionById(context.Background(), auctionId)

	if findAuctionError != nil {
		errRest := rest_err.ConvertError(findAuctionError)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}

func (ac *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
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

	auctionWinningData, findAuctionError := ac.AuctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)

	if findAuctionError != nil {
		errRest := rest_err.ConvertError(findAuctionError)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionWinningData)
}
