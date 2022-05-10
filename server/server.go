package server

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/victorcel/crud-rest-vozy/database"
	"github.com/victorcel/crud-rest-vozy/repository"
	"log"
	"net/http"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database is required")
	}

	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewMongoDbRepository(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal("Error server mongodb", err)
	}

	repository.SetUserRepository(repo)
	repository.SetPostRepository(repo)

	log.Println("Starting server on port", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
