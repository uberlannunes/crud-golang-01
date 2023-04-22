package setup

import (
	"github.com/uberlannunes/crud-golang-01/controller"
	"github.com/uberlannunes/crud-golang-01/model/repository"
	"github.com/uberlannunes/crud-golang-01/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {

	repo := repository.NewUserMongoDBRepository(database)
	userService := service.NewUserDomainService(repo)
	userController := controller.NewUserController(userService)

	return userController
}
