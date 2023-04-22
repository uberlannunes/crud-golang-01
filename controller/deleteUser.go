package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) DeleteUser(c *gin.Context) {

	logger.Info("call DeleteUserController...")

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError(fmt.Sprintf("'%s' is not a valid ID", userId))

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("jorney", "deleteUser"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"user deleted successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusOK)
}
