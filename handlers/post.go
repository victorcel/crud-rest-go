package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/victorcel/crud-rest-vozy/useCases/post"
	"net/http"
)

func InsertPost(validate *validator.Validate) http.HandlerFunc {
	return post.InsertPost(validate)
}

func GetPosts() http.HandlerFunc {
	return post.GetPosts()
}
