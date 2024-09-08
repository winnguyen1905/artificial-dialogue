package controller

import (
	"artificial-dialogue/internal/service"
	"artificial-dialogue/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(c *gin.Context);
	Hello();
}

type userController struct {
	userService service.UserService
}

func (uc *userController) Hello() {
    print("Hello!")
}

func (uc *userController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 200, "User")
}

func NewUserController() UserController {
	return &userController{
		userService: service.NewUserService(),
	}
}