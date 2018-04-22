package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/harkce/beasiswakita"
	"github.com/harkce/beasiswakita/server"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Starting beasiswakita...")
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/harkce/beasiswakita/.env")

	dsn := os.Getenv("DEVELOPMENT_DATABASE_URL")

	beasiswakita.InitDb(dsn)

	router := server.Router()

	log.Println("Beasiswakita started @:8061")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		http.ListenAndServe(":8061", router)
	}()

	<-sigChan
	log.Println("Shutting down beasiswakita...")
	log.Println("Beasiswakita stopped")
}
