package repository

import (
	"context"
	"os"

	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("call deleteUser repository")

	collection_users := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_users)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error(
			"Error trying to delete user",
			err,
			zap.String("userId", userId),
			zap.String("journey", "deleteUser"),
		)

		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"deleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
