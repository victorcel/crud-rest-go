package user

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/victorcel/crud-rest-vozy/models"
	"github.com/victorcel/crud-rest-vozy/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const HashCost = 10

type responseError struct {
	Message string
}

type InsertUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InsertUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InsertUser(validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var request = InsertUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HashCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user = models.User{
			ID:       primitive.NewObjectID(),
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashPassword),
		}

		err = validate.Struct(user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseError{
				Message: err.Error(),
			})
			return
		}

		email, errorEmail := repository.GetUserByEmail(r.Context(), user.Email)
		if errorEmail != nil && errorEmail.Error() != "mongo: no documents in result" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseError{
				Message: errorEmail.Error(),
			})
			return
		}

		if (email != models.User{}) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(responseError{
				Message: "Email ya existe",
			})

			return
		}

		repositoryInsertUser, err := repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(InsertUserResponse{
			ID:    repositoryInsertUser,
			Name:  user.Name,
			Email: user.Email,
		})

	}
}
