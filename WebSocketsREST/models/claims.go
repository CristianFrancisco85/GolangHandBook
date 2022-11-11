package models

import "github.com/golang-jwt/jwt"

// Se hace una composicion con la estructura jwt.StandardClaims para poder agregar campos personalizados
type AppClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
