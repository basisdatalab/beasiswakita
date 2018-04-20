package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/harkce/beasiswakita"
	"github.com/harkce/beasiswakita/server"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/harkce/beasiswakita/.env")

	dsn := os.Getenv("DEVELOPMENT_DATABASE_URL")

	beasiswakita.InitDb(dsn)

	router := server.Router()

	fmt.Println("Listening on port 8061")
	http.ListenAndServe(":8061", router)
}
