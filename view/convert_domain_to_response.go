package view

import (
	"github.com/uberlannunes/crud-golang-01/controller/model/response"
	"github.com/uberlannunes/crud-golang-01/model"
)

func ConvertUserDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}

}
