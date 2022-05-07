package database

import (
	"context"
	"crud-rest-vozy/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const CollectionName = "users"

type MongoDbRepository struct {
	db *mongo.Database
}

func NewMongoDbRepository(url string) (*MongoDbRepository, error) {

	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connect mongodb", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error ping mongodb", err)
	}

	db := client.Database("rest-vozy")
	return &MongoDbRepository{db}, nil
}

func (repository *MongoDbRepository) Close() error {
	return repository.db.Client().Disconnect(context.Background())
}

func (repository *MongoDbRepository) InsertUser(ctx context.Context, user *models.User) (string, error) {
	result, err := repository.db.Collection(CollectionName).InsertOne(ctx, user)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex()), err
}

func (repository *MongoDbRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {

	var user models.User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	err = repository.db.Collection(CollectionName).FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	var user models.User

	err := repository.db.Collection(CollectionName).FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var user models.User
	var users []models.User

	cursor, err := repository.db.Collection(CollectionName).Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return users, err
	}

	for cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository *MongoDbRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, name string) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}
	updateOne, err := repository.db.Collection(CollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}

func (repository *MongoDbRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	deleteOne, err := repository.db.Collection(CollectionName).DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return nil, err
	}
	return deleteOne, nil
}
