package user

import (
	"crud-rest-vozy/repository"
	"encoding/json"
	"net/http"
)

type updateUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type updateUserResponse struct {
	Message string `json:"message"`
}

func UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request = updateUserRequest{}

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updateUser, err := repository.UpdateUser(r.Context(), request.Id, request.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		var messageResponse = "user not exist"

		if updateUser.MatchedCount >= 1 {
			messageResponse = "update successful"
		}

		_ = json.NewEncoder(w).Encode(updateUserResponse{
			Message: messageResponse,
		})
	}
}
