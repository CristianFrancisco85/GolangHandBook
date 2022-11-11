package main

import (
	"context"
	"log"
	"os"
	"rest_ws/handlers"
	"rest_ws/middleware"
	"rest_ws/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Se carga el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Se crean las variables de entorno
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Se crea el servidor REST y Websockets
	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Se inicia el servidor REST y Websockets
	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {

	// Se crea un subrouter para las el apth que utiliza el middleware de autenticación
	api := r.PathPrefix("/api/v1").Subrouter()

	// Se le aplica el middleware de autenticación
	// A este middleware se le pasa el servidor para poder acceder a la configuración donde se encuentra la clave secreta
	api.Use(middleware.CheckAuthMiddleware(s))

	// Se crean las rutas que van sobre la raíz del servidor
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods("POST")

	// Se registran las rutas del middleware de autenticación
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods("GET")
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods("POST")
	api.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods("GET")
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods("PUT")
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods("DELETE")

	// Se registran las rutas de websockets
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)

}
