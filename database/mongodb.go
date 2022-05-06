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

var (
	userCollection *mongo.Collection
)

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
	userCollection = db.Collection("users")
	return &MongoDbRepository{db}, nil
}

func (repository *MongoDbRepository) Close() error {
	return repository.db.Client().Disconnect(context.Background())
}

func (repository *MongoDbRepository) InsertUser(ctx context.Context, user *models.User) (string, error) {
	result, err := userCollection.InsertOne(ctx, user)

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

	err = userCollection.FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	var user models.User

	err := userCollection.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var user models.User
	var users []models.User

	cursor, err := userCollection.Find(ctx, bson.D{})
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

func (repository *MongoDbRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, name string) error {
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}
	_, err := userCollection.UpdateOne(ctx, filter, update)
	return err
}

func (repository *MongoDbRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := userCollection.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
