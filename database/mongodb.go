package database

import (
	"context"
	"fmt"
	"github.com/victorcel/crud-rest-vozy/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	CollectionNameUser = "users"
	CollectionNamePost = "posts"
)

type MongoDbRepository struct {
	db *mongo.Database
}

type UserWithoutPassword struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
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
	result, err := repository.db.Collection(CollectionNameUser).InsertOne(ctx, user)

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

	err = repository.db.Collection(CollectionNameUser).FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	var user models.User

	err := repository.db.Collection(CollectionNameUser).FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repository *MongoDbRepository) GetUsers(ctx context.Context) ([]UserWithoutPassword, error) {
	var user UserWithoutPassword
	var users []UserWithoutPassword

	cursor, err := repository.db.Collection(CollectionNameUser).Find(ctx, bson.D{})
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

func (repository *MongoDbRepository) UpdateUser(ctx context.Context, id string, name string) (
	*mongo.UpdateResult, error,
) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"name", name}}}}
	updateOne, err := repository.db.Collection(CollectionNameUser).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}

func (repository *MongoDbRepository) DeleteUser(ctx context.Context, id string) (
	*mongo.DeleteResult, error,
) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	deleteOne, err := repository.db.Collection(CollectionNameUser).DeleteOne(ctx, bson.D{{"_id", objectId}})
	if err != nil {
		return nil, err
	}
	return deleteOne, nil
}

func (repository *MongoDbRepository) InsertPost(ctx context.Context, post *models.Post) (string, error) {
	result, err := repository.db.Collection(CollectionNamePost).InsertOne(ctx, post)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex()), err
}

func (repository *MongoDbRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return &post, err
	}

	err = repository.db.Collection(CollectionNamePost).FindOne(ctx, bson.D{{"_id", objectId}}).Decode(&post)
	if err != nil {
		return &models.Post{}, err
	}

	return &post, nil
}

func (repository *MongoDbRepository) DeletePost(ctx context.Context, id primitive.ObjectID) (
	*mongo.DeleteResult, error,
) {
	deleteOne, err := repository.db.Collection(CollectionNamePost).DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return nil, err
	}
	return deleteOne, nil
}

func (repository *MongoDbRepository) UpdatePost(ctx context.Context, post *models.Post) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id", post.Id}}
	update := bson.D{{"$set", bson.D{{"post_content", post.PostContent}}}}
	updateOne, err := repository.db.Collection(CollectionNamePost).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateOne, nil
}

func (repository *MongoDbRepository) GetPosts(ctx context.Context) (*[]models.Post, error) {
	var post models.Post
	var posts []models.Post

	cursor, err := repository.db.Collection(CollectionNamePost).Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return &posts, err
	}

	for cursor.Next(ctx) {
		err := cursor.Decode(&post)
		if err != nil {
			return &posts, err
		}
		posts = append(posts, post)
	}

	return &posts, nil
}
