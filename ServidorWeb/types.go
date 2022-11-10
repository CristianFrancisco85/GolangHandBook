package main

import (
	"net/http"
)

//Un middleware en un handler que recibe el handler que se ejecuta despues del middleware, este puede ser otro middleware o el handler final
type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
