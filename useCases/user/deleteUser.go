package user

import (
	"crud-rest-vozy/repository"
	"encoding/json"
	"net/http"
)

type deleteUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type deleteUserResponse struct {
	Message string `json:"message"`
}

func DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = deleteUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		deleteUser, err := repository.DeleteUser(r.Context(), request.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var messageResponse = "user not exist"

		if deleteUser.DeletedCount >= 1 {
			messageResponse = "delete successful"
		}

		_ = json.NewEncoder(w).Encode(deleteUserResponse{
			Message: messageResponse,
		})
	}
}
