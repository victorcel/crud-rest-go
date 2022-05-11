package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	Id          primitive.ObjectID `bson:"_id"`
	PostContent string             `bson:"post_content" validate:"required"`
	CreatedAt   time.Time          `bson:"createdAt" validate:"required"`
	UserId      string             `bson:"userId" validate:"required"`
}
