package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/uberlannunes/crud-golang-01/controller"
)

func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
