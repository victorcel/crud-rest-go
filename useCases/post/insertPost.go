package post

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/victorcel/crud-rest-vozy/models"
	"github.com/victorcel/crud-rest-vozy/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type responseError struct {
	Message string
}

type InsertPostRequest struct {
	Id          primitive.ObjectID `json:"_id"`
	PostContent string             `json:"post_content"`
	CreatedAt   time.Time          `json:"createdAt"`
	UserId      string             `json:"user_id"`
}

type InsertPostResponse struct {
	Id          primitive.ObjectID `json:"_id"`
	PostContent string             `json:"post_content"`
	CreatedAt   time.Time          `json:"createdAt"`
	UserId      string             `json:"user_id"`
}

func InsertPost(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var request = InsertPostRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		getPostByID, err := repository.GetUserByID(r.Context(), request.UserId)

		if err != nil && getPostByID.ID.IsZero() {
			http.Error(w, "El usuario no existe", http.StatusBadRequest)
			return
		}

		var post = models.Post{
			Id:          primitive.NewObjectID(),
			PostContent: request.PostContent,
			CreatedAt:   time.Now(),
			UserId:      request.UserId,
		}

		err = validate.Struct(post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseError{
				Message: err.Error(),
			})
			return
		}

		repositoryInsertUser, err := repository.InsertPost(r.Context(), &post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		objectId, err := primitive.ObjectIDFromHex(repositoryInsertUser)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(InsertPostResponse{
			Id:          objectId,
			PostContent: post.PostContent,
			UserId:      post.UserId,
			CreatedAt:   post.CreatedAt,
		})

	}
}
