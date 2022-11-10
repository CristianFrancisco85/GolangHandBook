package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {

	//Se recibe el handler que va despues para poder crear el nuevo handler con el middleware aplicado
	return func(f http.HandlerFunc) http.HandlerFunc {

		//Retorna el handler que ejecuta la logica del middleware y luego ejecuta el handler que le sigue
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication ...")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	//Se recibe el handler que va despues para poder crear el nuevo handler con el middleware aplicado
	return func(f http.HandlerFunc) http.HandlerFunc {

		//Retorna el handler que ejecuta la logica del middleware y luego ejecuta el handler que le sigue
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
