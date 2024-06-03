package user_entity

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

type UserRepositoryInterface interface {
	FindUserById(ctx context.Context, id string) (*User, *internal_error.InternalError)
}
