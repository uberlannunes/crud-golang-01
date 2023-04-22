package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Passsword string             `bson:"password,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Age       int                `bson:"age,omitempty"`
}
