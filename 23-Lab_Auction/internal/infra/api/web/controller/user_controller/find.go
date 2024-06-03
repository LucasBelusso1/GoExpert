package user_controller

import (
	"context"
	"net/http"

	"github.com/LucasBelusso1/23-Lab_Auction/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (u *UserController) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	err := uuid.Validate(userId)

	if err != nil {
		errRest := rest_err.NewBadRequestError("InvalidFields", rest_err.Causes{
			Field:   "userId",
			Message: "Invalid UUID value",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	userData, findUserError := u.UserUseCase.FindUserById(context.Background(), userId)

	if findUserError != nil {
		errRest := rest_err.ConvertError(findUserError)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, userData)
}
