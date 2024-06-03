package user_usecase

import user_entity "github.com/LucasBelusso1/23-Lab_Auction/internal/entity/user"

type UserUseCase struct {
	UserRepository user_entity.UserRepositoryInterface
}

func NewUserUseCase(userRepository user_entity.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{UserRepository: userRepository}
}
