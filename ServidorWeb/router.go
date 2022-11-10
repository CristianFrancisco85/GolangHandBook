package main

import (
	"net/http"
)

//Struct de nuestro router, tiene un map donde el path es la key y el handler el valor
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//Constructor del router
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//Funcion que nos dice si existe un handler para cierta ruta
func (r *Router) FindHandler(method, path string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, pathExist
}

//Funcion que sirve como punto de entrada, si no existe la ruta regresa un 404, y si existe ejecuta el handler
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(request.Method, request.URL.Path)
	if !pathExist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
