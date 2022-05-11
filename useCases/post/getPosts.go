package post

import (
	"encoding/json"
	"github.com/victorcel/crud-rest-vozy/repository"
	"net/http"
)

func GetPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		listPost, err := repository.GetPosts(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(listPost)
	}
}
