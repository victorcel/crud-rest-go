package handlers

import (
	"crud-rest-vozy/useCases/user"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func InsertUser(validate *validator.Validate) http.HandlerFunc {
	return user.InsertUser(validate)
}

func GetUserByID() http.HandlerFunc {
	return user.GetUserByID()
}

func GetUsers() http.HandlerFunc {
	return user.GetUsers()
}

func UpdateUser() http.HandlerFunc {
	return user.UpdateUser()
}

func DeleteUser() http.HandlerFunc {
	return user.DeleteUser()
}
