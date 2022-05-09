package user

import (
	"crud-rest-vozy/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type IdRequest struct {
	Id string `json:"id"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		id := params["id"]
		getUser, err := repository.GetUserByID(r.Context(), id)

		if err != nil && err.Error() == "the provided hex string is not a valid ObjectID" {
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, fmt.Sprintf("Id %s no existe", id), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(GetUserResponse{
			ID:    getUser.ID.Hex(),
			Name:  getUser.Name,
			Email: getUser.Email,
		})
	}
}
