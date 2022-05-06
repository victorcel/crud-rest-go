package handlers

import (
	"crud-rest-vozy/models"
	"crud-rest-vozy/repository"
	"crud-rest-vozy/server"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const HashCost = 10

type SignUpLoginRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ResponseError struct {
	Message string
}

func SignUpHandler(s server.Server, validate *validator.Validate) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var request = SignUpLoginRequest{}

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
			json.NewEncoder(w).Encode(ResponseError{
				Message: err.Error(),
			})
			return
		}

		email, errorEmail := repository.GetUserByEmail(r.Context(), user.Email)
		if errorEmail != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseError{
				Message: errorEmail.Error(),
			})

			return
		}

		if (email != models.User{}) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ResponseError{
				Message: "Email ya existe",
			})

			return
		}

		repositoryInsertUser, err := repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(SignUpResponse{
			ID:    repositoryInsertUser,
			Name:  user.Name,
			Email: user.Email,
		})

	}
}
