package service

import (
	"github.com/uberlannunes/crud-golang-01/configuration/logger"
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("call CreateUserService", zap.String("jorney", "createUser"))

	user, _ := ud.FindUserByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already registed")
	}

	userDomain.EncryptPassword()
	userCreated, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
