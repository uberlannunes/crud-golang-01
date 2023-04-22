package repository

import (
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

type UserRepositoryInterface interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserMongoDBRepository(dbConnection *mongo.Database) UserRepositoryInterface {
	return &userRepository{
		databaseConnection: dbConnection,
	}
}
