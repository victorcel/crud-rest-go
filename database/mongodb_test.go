package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/victorcel/crud-rest-vozy/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func setup() *mongo.Database {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URL"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connect mongodb", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error ping mongodb", err)
	}
	return client.Database(os.Getenv("DATABASE_NAME_TESTING"))
}

func TestMongoDbRepository_Close(t *testing.T) {
	type fields struct {
		db *mongo.Database
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Close",
			fields:  fields{db: setup()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			if err := repository.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMongoDbRepository_GetUserByEmail(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     "victor",
		Email:    "victor@local.com",
		Password: "23233",
	}

	_, err := repository.Collection(CollectionNameUser).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "GetUserByEmail",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx:   context.Background(),
				email: user.Email,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Error - GetUserByEmail",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx:   context.Background(),
				email: "user@lical.com",
			},
			want:    models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.GetUserByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDbRepository_GetUserByID(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     "victor",
		Email:    "victor@local.com",
		Password: "23233",
	}

	_, err := repository.Collection(CollectionNameUser).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "GetUserByID",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				id:  user.ID.Hex(),
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "Error - GetUserByID",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				id:  primitive.NewObjectID().Hex(),
			},
			want:    models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.GetUserByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDbRepository_GetUsers(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	user1 := UserWithoutPassword{
		Id:    primitive.NewObjectID(),
		Name:  "victor",
		Email: "victor@local.com",
	}
	user2 := UserWithoutPassword{
		Id:    primitive.NewObjectID(),
		Name:  "elias",
		Email: "elias@local.com",
	}
	users := []interface{}{user1, user2}

	_, err := repository.Collection(CollectionNameUser).InsertMany(context.Background(), users)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []UserWithoutPassword
		wantErr bool
	}{
		{
			name: "GetUsers",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    []UserWithoutPassword{user1, user2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.GetUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDbRepository_InsertUser(t *testing.T) {
	objectID := primitive.NewObjectID()
	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx  context.Context
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "InsertUser",
			fields: fields{db: setup()},
			args: args{
				ctx: context.Background(),
				user: &models.User{
					ID:       objectID,
					Name:     "victor",
					Email:    "vi@local.com",
					Password: "23232",
				},
			},
			want:    objectID.Hex(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.InsertUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InsertUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoDbRepository_UpdateUser(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     "victor",
		Email:    "victor@local.com",
		Password: "23233",
	}

	_, err := repository.Collection(CollectionNameUser).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx  context.Context
		id   string
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "UpdateUser",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx:  context.Background(),
				id:   user.ID.Hex(),
				name: "elias",
			},
			want: &mongo.UpdateResult{
				MatchedCount:  1,
				ModifiedCount: 1,
				UpsertedCount: 0,
				UpsertedID:    nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "No existe ID - UpdateUser",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx:  context.Background(),
				id:   primitive.NewObjectID().Hex(),
				name: "elias",
			},
			want: &mongo.UpdateResult{
				MatchedCount:  0,
				ModifiedCount: 0,
				UpsertedCount: 0,
				UpsertedID:    nil,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.UpdateUser(tt.args.ctx, tt.args.id, tt.args.name)
			if !tt.wantErr(t, err, fmt.Sprintf("UpdateUser(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.name)) {
				return
			}
			assert.Equalf(t, tt.want, got, "UpdateUser(%v, %v, %v)", tt.args.ctx, tt.args.id, tt.args.name)
		})
	}
}

func TestMongoDbRepository_DeleteUser(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	user1 := models.User{
		ID:       primitive.NewObjectID(),
		Name:     "victor",
		Email:    "victor@local.com",
		Password: "23233",
	}

	_, err := repository.Collection(CollectionNameUser).InsertOne(context.Background(), user1)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "DeleteUser",
			fields: fields{db: setup()},
			args: args{
				ctx: context.Background(),
				id:  user1.ID.Hex(),
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantErr: assert.NoError,
		},
		{
			name:   "Error - DeleteUser",
			fields: fields{db: setup()},
			args: args{
				ctx: context.Background(),
				id:  primitive.NewObjectID().Hex(),
			},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.DeleteUser(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("DeleteUser(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "DeleteUser(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func TestNewMongoDbRepository(t *testing.T) {
	assert.NotEmpty(t, setup())
}

func TestMongoDbRepository_InsertPost(t *testing.T) {
	objectID := primitive.NewObjectID()

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx  context.Context
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "InsertPost",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				post: &models.Post{
					Id:          objectID,
					PostContent: "Cafe internet",
					CreatedAt:   time.Now(),
					UserId:      primitive.NewObjectID().Hex(),
				},
			},
			want:    objectID.Hex(),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.InsertPost(tt.args.ctx, tt.args.post)
			if !tt.wantErr(t, err, fmt.Sprintf("InsertPost(%v, %v)", tt.args.ctx, tt.args.post)) {
				return
			}
			assert.Equalf(t, tt.want, got, "InsertPost(%v, %v)", tt.args.ctx, tt.args.post)
		})
	}
}

func TestMongoDbRepository_GetPostByID(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	post := models.Post{
		Id:          primitive.NewObjectID(),
		PostContent: "Cafe internet",
		CreatedAt:   time.Time{},
		UserId:      primitive.NewObjectID().Hex(),
	}

	_, err := repository.Collection(CollectionNamePost).InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}
	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Post
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "GetPostByID",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				id:  post.Id.Hex(),
			},
			want:    &post,
			wantErr: assert.NoError,
		},
		{
			name: "Error - GetPostByID",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				id:  primitive.NewObjectID().Hex(),
			},
			want:    &models.Post{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.GetPostByID(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetPostByID(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetPostByID(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func TestMongoDbRepository_DeletePost(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	post := models.Post{
		Id:          primitive.NewObjectID(),
		PostContent: "Cafe internet",
		CreatedAt:   time.Time{},
		UserId:      primitive.NewObjectID().Hex(),
	}

	_, err := repository.Collection(CollectionNamePost).InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
		id  primitive.ObjectID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mongo.DeleteResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "DeletePost",
			fields: fields{db: setup()},
			args: args{
				ctx: context.Background(),
				id:  post.Id,
			},
			want: &mongo.DeleteResult{
				DeletedCount: 1,
			},
			wantErr: assert.NoError,
		},
		{
			name:   "Error - DeletePost",
			fields: fields{db: setup()},
			args: args{
				ctx: context.Background(),
				id:  primitive.NewObjectID(),
			},
			want: &mongo.DeleteResult{
				DeletedCount: 0,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.DeletePost(tt.args.ctx, tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("DeletePost(%v, %v)", tt.args.ctx, tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "DeletePost(%v, %v)", tt.args.ctx, tt.args.id)
		})
	}
}

func TestMongoDbRepository_UpdatePost(t *testing.T) {
	_ = setup().Drop(context.Background())
	repository := setup()
	post := models.Post{
		Id:          primitive.NewObjectID(),
		PostContent: "Cafe internet",
		CreatedAt:   time.Time{},
		UserId:      primitive.NewObjectID().Hex(),
	}

	_, err := repository.Collection(CollectionNamePost).InsertOne(context.Background(), post)
	if err != nil {
		log.Fatal(err)
	}

	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx  context.Context
		post *models.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *mongo.UpdateResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "UpdatePost",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx: context.Background(),
				post: &models.Post{
					Id:          post.Id,
					PostContent: "Cafe Playa",
				},
			},
			want: &mongo.UpdateResult{
				MatchedCount:  1,
				ModifiedCount: 1,
				UpsertedCount: 0,
				UpsertedID:    nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "No existe ID - UpdatePost",
			fields: fields{
				db: setup(),
			},
			args: args{
				ctx:  context.Background(),
				post: &models.Post{},
			},
			want: &mongo.UpdateResult{
				MatchedCount:  0,
				ModifiedCount: 0,
				UpsertedCount: 0,
				UpsertedID:    nil,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.UpdatePost(tt.args.ctx, tt.args.post)
			if !tt.wantErr(t, err, fmt.Sprintf("UpdatePost(%v, %v)", tt.args.ctx, tt.args.post)) {
				return
			}
			assert.Equalf(t, tt.want, got, "UpdatePost(%v, %v)", tt.args.ctx, tt.args.post)
		})
	}
}

func TestMongoDbRepository_ListPost(t *testing.T) {
	type fields struct {
		db *mongo.Database
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *[]models.Post
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &MongoDbRepository{
				db: tt.fields.db,
			}
			got, err := repository.ListPost(tt.args.ctx)
			if !tt.wantErr(t, err, fmt.Sprintf("ListPost(%v)", tt.args.ctx)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ListPost(%v)", tt.args.ctx)
		})
	}
}
