package repository

import (
	"context"
	"crud-rest-vozy/database"
	"crud-rest-vozy/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) (string, error)
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUsers(ctx context.Context) ([]database.UserWithoutPassword, error)
	UpdateUser(ctx context.Context, id string, name string) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error)
	Close() error
}

var implementationUser UserRepository

func SetUserRepository(repository UserRepository) {
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

func GetUsers(ctx context.Context) ([]database.UserWithoutPassword, error) {
	return implementationUser.GetUsers(ctx)
}

func UpdateUser(ctx context.Context, id string, name string) (*mongo.UpdateResult, error) {
	return implementationUser.UpdateUser(ctx, id, name)
}

func DeleteUser(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return implementationUser.DeleteUser(ctx, id)
}

func Close() error {
	return implementationUser.Close()
}
