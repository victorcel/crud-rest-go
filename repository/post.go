package repository

import (
	"context"
	"crud-rest-vozy/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository interface {
	InsertPost(ctx context.Context, post *models.Post) (string, error)
	GetPostByID(ctx context.Context, id string) (*models.Post, error)
	DeletePost(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	UpdatePost(ctx context.Context, post *models.Post) (*mongo.UpdateResult, error)
	ListPost(ctx context.Context) (*[]models.Post, error)
}

var implementationPost PostRepository

func SetPostRepository(repository PostRepository) {
	implementationPost = repository
}

func InsertPost(ctx context.Context, post *models.Post) (string, error) {
	return implementationPost.InsertPost(ctx, post)
}

func GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	return implementationPost.GetPostByID(ctx, id)
}

func DeletePost(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return implementationPost.DeletePost(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) (*mongo.UpdateResult, error) {
	return implementationPost.UpdatePost(ctx, post)
}

func ListPost(ctx context.Context) (*[]models.Post, error) {
	return implementationPost.ListPost(ctx)
}
