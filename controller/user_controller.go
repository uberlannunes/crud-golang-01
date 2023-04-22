package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/uberlannunes/crud-golang-01/model/service"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userController struct {
	service service.UserDomainServiceInterface
}

func NewUserController(
	userService service.UserDomainServiceInterface,
) UserControllerInterface {
	return &userController{
		service: userService,
	}
}
