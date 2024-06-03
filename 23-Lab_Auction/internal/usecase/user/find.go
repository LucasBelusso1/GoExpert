package user_usecase

import (
	"context"

	"github.com/LucasBelusso1/23-Lab_Auction/internal/internal_error"
)

func (u *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internal_error.InternalError) {
	userEntity, err := u.UserRepository.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		Id:   userEntity.Id,
		Name: userEntity.Name,
	}, nil
}
