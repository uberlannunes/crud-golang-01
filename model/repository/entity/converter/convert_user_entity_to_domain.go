package converter

import (
	"github.com/uberlannunes/crud-golang-01/model"
	"github.com/uberlannunes/crud-golang-01/model/repository/entity"
)

func ConvertUserEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Passsword,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
