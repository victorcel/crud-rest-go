package repository

import (
	"context"
	"crud-rest-vozy/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User interface {
	InsertUser(ctx context.Context, user *models.User) (string, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, name string) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	Close() error
}

var implementationUser User

func SetRepository(repository User) {
	implementationUser = repository
}

func InsertUser(ctx context.Context, user *models.User) (string, error) {
	return implementationUser.InsertUser(ctx, user)
}

func GetUserByID(ctx context.Context, id string) (models.User, error) {
	return implementationUser.GetUserByID(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	return implementationUser.GetUserByEmail(ctx, email)
}

func GetUsers(ctx context.Context) ([]models.User, error) {
	return implementationUser.GetUsers(ctx)
}

func UpdateUser(ctx context.Context, id primitive.ObjectID, name string) error {
	return implementationUser.UpdateUser(ctx, id, name)
}

func DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return implementationUser.DeleteUser(ctx, id)
}

func Close() error {
	return implementationUser.Close()
}
