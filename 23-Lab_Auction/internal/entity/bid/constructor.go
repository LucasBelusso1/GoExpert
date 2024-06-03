package bid_entity

import (
	"time"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	"github.com/google/uuid"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

func NewBidEntity(userId, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.NewString(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	err := bid.Validate()

	if err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	err := uuid.Validate(b.UserId)

	if err != nil {
		return internal_error.NewBadRequestError("UserId is not a valid ID")
	}

	err = uuid.Validate(b.AuctionId)

	if err != nil {
		return internal_error.NewBadRequestError("AuctionId is not a valid ID")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount is not a valid")
	}

	return nil
}
