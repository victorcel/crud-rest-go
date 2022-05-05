package models

import "time"

type Post struct {
	Id          string    `bson:"id"`
	PostContent string    `bson:"postContent"`
	CreatedAt   time.Time `bson:"createdAt"`
	UserId      string    `bson:"userId"`
}
