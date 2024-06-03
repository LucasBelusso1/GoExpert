package auction_entity

import (
	"time"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
	"github.com/google/uuid"
)

type ProductCondition int

const (
	New ProductCondition = iota
	Used
	Refurbished
)

type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

func NewAuction(productName, cateogry, description string, condition ProductCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.NewString(),
		ProductName: productName,
		Category:    cateogry,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	err := auction.Validate()

	if err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() *internal_error.InternalError {
	if len(a.ProductName) <= 1 {
		return internal_error.NewBadRequestError("Product name must contain more than one character")
	}

	if len(a.Category) <= 2 {
		return internal_error.NewBadRequestError("Category must contain more than two characters")
	}

	if len(a.Description) < 10 {
		return internal_error.NewBadRequestError("Description must contain at least 10 characters")
	}

	if a.Condition != New && a.Condition != Used && a.Condition != Refurbished {
		return internal_error.NewBadRequestError("Invalid condition status")
	}

	return nil
}
