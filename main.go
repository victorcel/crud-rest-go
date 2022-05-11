package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/victorcel/crud-rest-vozy/handlers"
	"github.com/victorcel/crud-rest-vozy/server"
	"log"
	"net/http"
	"os"
)

var validate *validator.Validate

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := server.NewServer(&server.Config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	})

	if err != nil {
		log.Fatalf("Error creating server %v\n", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	validate = validator.New()
	r.HandleFunc("/api/v1/user", handlers.InsertUser(validate)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users", handlers.GetUsers()).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/{id}", handlers.GetUserByID()).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user", handlers.UpdateUser()).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/user", handlers.DeleteUser()).Methods(http.MethodDelete)

	r.HandleFunc("/api/v1/post", handlers.InsertPost(validate)).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/posts", handlers.GetPosts()).Methods(http.MethodGet)

}
