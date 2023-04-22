package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/configuration/validation"
	"github.com/uberlannunes/crud-golang-01/controller/model/request"
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userController) UpdateUser(c *gin.Context) {

	logger.Info("call UpdateUserController...")

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := rest_err.NewBadRequestError(fmt.Sprintf("'%s' is not a valid ID", userId))

		logger.Error(
			errorMessage.Message,
			err,
			zap.String("jorney", "updateUser"),
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user data", err, zap.String("journey", "updateUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	request := view.ConvertUserDomainToResponse(domain)

	logger.Info(
		fmt.Sprintf("user updated successfully, request=%v", request),
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusOK)
}
