package service

import (
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(userId string) *rest_err.RestErr {

	logger.Info("call DeleteUser service", zap.String("jorney", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error(
			"Error trying to call deleteUser repository",
			err,
			zap.String("userId", userId),
			zap.String("journey", "deleteUser"),
		)
		return err
	}

	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
