package server

import (
	"crud-rest-vozy/database"
	"crud-rest-vozy/repository"
	"errors"
	"github.com/gorilla/mux"
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

	repository.SetRepository(repo)

	log.Println("Starting server on port", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
