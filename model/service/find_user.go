package service

import (
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByID(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByID service", zap.String("jorney", "findUserById"))
	return ud.userRepository.FindUserByID(userId)
}

func (ud *userDomainService) FindUserByEmail(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("init FindUserByEmail service", zap.String("jorney", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(userEmail)
}
