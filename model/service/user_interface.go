package service

import (
	"github.com/uberlannunes/crud-golang-01/configuration/rest_err"
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/model/repository"
)

type UserDomainServiceInterface interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByID(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserDomainService(
	repository repository.UserRepositoryInterface,
) UserDomainServiceInterface {
	return &userDomainService{
		userRepository: repository,
	}
}
