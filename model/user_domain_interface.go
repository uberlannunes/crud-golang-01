package model

type UserDomainInterface interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int
	GetJSONValue() (string, error)
	EncryptPassword()

	SetID(string)
}
