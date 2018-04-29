package file

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/basisdatalab/beasiswakita/authentication"
	"github.com/basisdatalab/beasiswakita/errors"
	"github.com/basisdatalab/beasiswakita/server/response"
	"github.com/julienschmidt/httprouter"
)

type FileHandler struct{}

func (h *FileHandler) Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := authentication.Token(r.Header.Get("Authorization"), []string{})
	if err != nil {
		response.Error(w, errors.Unauthorized)
		log.Println(err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		response.Error(w, errors.BadRequest)
		log.Println(err)
		return
	}
	defer file.Close()

	extension := strings.Join(strings.Split(handler.Filename, ".")[1:], "")
	if !extensionAllowed(extension) {
		response.Error(w, errors.BadRequest)
		return
	}

	err = saveFile(handler.Filename, file)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}

	filePath := os.Getenv("BEASISWAKITA_HOST") + "/public/" + handler.Filename
	data := struct {
		URL string `json:"url"`
	}{
		URL: filePath,
	}

	response.OK(w, data)
	return
}

func (h *FileHandler) Retrieve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filename := ps.ByName("filename")
	path := os.Getenv("GOPATH") + "/src/github.com/basisdatalab/beasiswakita/public/" + filename

	if _, err := os.Stat(path); os.IsNotExist(err) {
		response.Error(w, errors.NotFound)
		log.Println(err)
		return
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		response.Error(w, errors.InternalServerError)
		log.Println(err)
		return
	}
	defer f.Close()

	response.File(w, f, path)
	return
}

func extensionAllowed(extension string) bool {
	extension = strings.ToLower(extension)
	allowedExtension := []string{"jpg", "jpeg", "png"}
	for _, ext := range allowedExtension {
		if ext == extension {
			return true
		}
	}
	return false
}

func saveFile(filename string, file io.Reader) error {
	path := os.Getenv("GOPATH") + "/src/github.com/basisdatalab/beasiswakita/public/" + filename
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, file)
	return nil
}
