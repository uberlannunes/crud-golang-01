package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) FindUserById(c *gin.Context) {
	logger.Info("init FindUserById controller", zap.String("jorney", "findUserById"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError(fmt.Sprintf("'%s' is not a valid ID", userId))

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("jorney", "findUserById"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByID(userId)
	if err != nil {
		errorMessage := rest_err.NewBadRequestError("Error trying to find User by ID")

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("userId", userId),
			zap.String("jorney", "findUserById"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response := view.ConvertUserDomainToResponse(userDomain)

	logger.Info(
		fmt.Sprintf("FindUserById executed successfully, response=%v", response),
		zap.String("userId", userId),
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, response)
}

func (uc *userController) FindUserByEmail(c *gin.Context) {
	logger.Info("init FindUserByEmail controller", zap.String("jorney", "findUserByEmail"))

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := rest_err.NewBadRequestError(fmt.Sprintf("'%s' is not a valid Email", userEmail))

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("jorney", "findUserByEmail"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmail(userEmail)
	if err != nil {
		errorMessage := rest_err.NewBadRequestError("Error trying to find User by ID")

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("userEmail", userEmail),
			zap.String("jorney", "findUserByEmail"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	response := view.ConvertUserDomainToResponse(userDomain)

	logger.Info(
		fmt.Sprintf("FindUserByEmail executed successfully, response=%v", response),
		zap.String("userEmail", userEmail),
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, response)
}
