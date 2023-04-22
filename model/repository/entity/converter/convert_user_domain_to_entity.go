package converter

import (
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/model/repository/entity"
)

func ConvertUserDomainToEntity(
	model model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:     model.GetEmail(),
		Passsword: model.GetPassword(),
		Name:      model.GetName(),
		Age:       model.GetAge(),
	}
}
