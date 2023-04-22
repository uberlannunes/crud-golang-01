package repository

import (
	"context"
	"os"

	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("call CreateUser repository")

	collection_users := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_users)

	userEntity := converter.ConvertUserDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), userEntity)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userEntity.ID = result.InsertedID.(primitive.ObjectID)
	return converter.ConvertUserEntityToDomain(*userEntity), nil
}
