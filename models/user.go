package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" validate:"required"`
	Name     string             `bson:"name" validate:"required"`
	Email    string             `bson:"email" validate:"required,email"`
	Password string             `bson:"password" validate:"required"`
}
