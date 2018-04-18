package router

import (
	"net/http"

	"github.com/goware/cors"
	"github.com/julienschmidt/httprouter"
)

func Router() http.Handler {
	router := httprouter.New()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	return cors.Handler(router)
}
