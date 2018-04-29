package server

import (
	"net/http"
	"os"

	"github.com/basisdatalab/beasiswakita/board"
	"github.com/basisdatalab/beasiswakita/file"
	"github.com/basisdatalab/beasiswakita/scholarship"
	"github.com/basisdatalab/beasiswakita/user"

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
	router.GET("/me", userHandler.Me)

	boardHandler := board.BoardHandler{}
	router.POST("/boards", boardHandler.Create)
	router.PUT("/boards", boardHandler.Update)
	router.PATCH("/boards/:boardID/state", boardHandler.State)
	router.DELETE("/boards", boardHandler.Delete)
	router.GET("/boards", boardHandler.Get)

	scholarshipHandler := scholarship.ScholarshipHandler{}
	router.POST("/scholarships", scholarshipHandler.Create)
	router.GET("/scholarships", scholarshipHandler.GetAll)
	router.GET("/scholarships/:scholarshipID", scholarshipHandler.Get)
	router.PUT("/scholarships", scholarshipHandler.Update)
	router.PATCH("/scholarships/:scholarshipID/state", scholarshipHandler.State)
	router.POST("/scholarships/:scholarshipID/whislist", scholarshipHandler.AddWhislist)

	fileHandler := file.FileHandler{}
	public := os.Getenv("GOPATH") + "/src/github.com/basisdatalab/beasiswakita/public"
	router.POST("/uploads", fileHandler.Upload)
	router.ServeFiles("/public/*filepath", http.Dir(public))

	return cors.Handler(router)
}
