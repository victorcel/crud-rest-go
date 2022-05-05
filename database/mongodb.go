package database

import (
	"context"
	"crud-rest-vozy/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Ctx = context.TODO()

type MongoDbRepository struct {
	db *mongo.Database
}

func NewMongoDbRepository(url string) (*MongoDbRepository, error) {

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("gallery-vozy")
	return &MongoDbRepository{db}, nil
}

func (repository *MongoDbRepository) Close() {
	err := repository.db.Client().Disconnect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (repository *MongoDbRepository) InsertUser(ctx context.Context, user *models.User) (string, error) {
	result, err := repository.db.Collection("users").InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex()), err
}

func (repository MongoDbRepository) name() {

}
