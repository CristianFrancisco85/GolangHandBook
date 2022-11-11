package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"rest_ws/database"
	"rest_ws/repository"
	"rest_ws/websockets"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port        string // Puerto en el que se va a ejecutar el servidor
	JWTSecret   string // Clave secreta para la generación de tokens
	DatabaseUrl string // Url de la base de datos
}

type Server interface {
	Config() *Config      // Devuelve la configuración del servidor
	Hub() *websockets.Hub // Devuelve el hub de websockets
}

// EL broker es la implementación del servidor
type Broker struct {
	config *Config
	router *mux.Router
	hub    *websockets.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websockets.Hub {
	return b.hub
}

// Crea un nuevo servidor y valida la configuración
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("JWT secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websockets.NewHub(),
	}

	return broker, nil

}

// Inicializa el broker del servidor
func (b *Broker) Start(binder func(s Server, r *mux.Router)) {

	// Se crea el router
	b.router = mux.NewRouter()
	// Se relaciona el router con el broker
	binder(b, b.router)
	// Se habilita el CORS
	handler := cors.Default().Handler(b.router)

	// Aqui se registra la implmentación específica de la base de datos
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	// A la implmentacion general del repositorio se le asigna la implementación específica
	repository.SetRepository(repo)

	// Se inicia el hub de websockets
	go b.hub.Run()

	// Se inicia el servidor
	log.Println("Server started on port", b.Config().Port)
	if err := http.ListenAndServe(":"+b.Config().Port, handler); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
