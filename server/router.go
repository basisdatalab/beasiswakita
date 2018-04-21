package server

import (
	"net/http"

	"github.com/harkce/beasiswakita/student_board"
	"github.com/harkce/beasiswakita/user"

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

	userHandler := user.UserHandler{}
	router.POST("/users", userHandler.Register)
	router.POST("/login", userHandler.Login)

	boardHandler := student_board.BoardHandler{}
	router.POST("/boards", boardHandler.Create)
	router.PUT("/boards", boardHandler.Update)
	router.PATCH("/boards/:boardID/state", boardHandler.State)

	return cors.Handler(router)
}
