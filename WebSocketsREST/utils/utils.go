package utils

import (
	"errors"
	"net/http"
	"rest_ws/models"
	"rest_ws/repository"
	"strings"

	"github.com/golang-jwt/jwt"
)

// Esta función se encarga de obtener el token de la cabecera de la petición y validarlo
func GetTokenFromHeader(r *http.Request, secret string) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

	// Se parsea el token
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// Esta función se encarga de obtener el los datos del usuario en función de las claims del token
func GetUserIdFromToken(r *http.Request, token *jwt.Token) (*models.User, error) {

	// Se valida las claims del token
	if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
		user, err := repository.FindUserById(r.Context(), claims.UserId)
		return user, err
	} else {
		return nil, errors.New("invalid token")
	}

}
