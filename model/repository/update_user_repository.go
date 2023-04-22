package repository

import (
	"context"
	"os"

	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("call UpdateUser repository")

	collection_users := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_users)

	userEntity := converter.ConvertUserDomainToEntity(userDomain)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	updateValue := bson.D{{Key: "$set", Value: userEntity}}

	_, err := collection.UpdateOne(context.Background(), filter, updateValue)
	if err != nil {
		logger.Error(
			"Error trying to update user",
			err,
			zap.String("userId", userId),
			zap.String("journey", "updateUser"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"UpdateUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	return nil
}
