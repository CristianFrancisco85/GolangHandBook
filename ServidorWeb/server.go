package main

import (
	"fmt"
	"net/http"
)

// Struct de nuestro server
type Server struct {
	port   string
	router *Router
}

//Constructor de nuestro server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Handle sirve para indicarle al router que cierta ruta ejecuta cierto handler
func (s *Server) Handle(method, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler

}

// Funcion variadica que retorna un handler que tiene los middlewares aplicados
func (s *Server) AddMiddleWare(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}

//funcion que inicia nuestro servidor
func (s *Server) Listen() error {

	//Se indica la ruta root y el router a donde redirige las peticiones
	http.Handle("/", s.router)

	//Se inicia y si hay un error se reporta
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
