package user_controller

import user_usecase "github.com/LucasBelusso1/23-Lab_Auction/internal/usecase/user"

type UserController struct {
	UserUseCase user_usecase.UserUseCaseInterface
}

func NewUserController(userUseCase user_usecase.UserUseCaseInterface) *UserController {
	return &UserController{UserUseCase: userUseCase}
}
