package response

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/harkce/beasiswakita/errors"
)

type ErrorInfo struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type Body struct {
	Data    interface{} `json:"data,omitempty"`
	Errors  []ErrorInfo `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
	Meta    Meta        `json:"meta"`
}

type Meta struct {
	HTTPStatus int `json:"http_status"`
	Limit      int `json:"limit,omitempty"`
	Offset     int `json:"offset,omitempty"`
	Total      int `json:"total,omitempty"`
}

func Send(status int, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(&Body{
		Data: data,
		Meta: Meta{
			HTTPStatus: status,
		},
	})
}

func SendMeta(status int, w http.ResponseWriter, data interface{}, meta Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(&Body{
		Data: data,
		Meta: Meta{
			HTTPStatus: status,
			Limit:      meta.Limit,
			Offset:     meta.Offset,
			Total:      meta.Total,
		},
	})
}

func OK(w http.ResponseWriter, data interface{}) {
	Send(200, w, data)
}

func OKMeta(w http.ResponseWriter, data interface{}, meta Meta) {
	SendMeta(200, w, data, meta)
}

func Created(w http.ResponseWriter, data interface{}) {
	Send(201, w, data)
}

func Accepted(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)

	json.NewEncoder(w).Encode(&Body{
		Message: message,
		Meta: Meta{
			HTTPStatus: 202,
		},
	})
}

func File(w http.ResponseWriter, file io.Reader, path string) {
	filename := filepath.Base(path)
	ext := filepath.Ext(path)
	mime := mime.TypeByExtension(ext)

	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	w.Header().Set("Content-Type", mime)
	w.WriteHeader(200)
	io.Copy(w, file)
}

func SendError(status int, w http.ResponseWriter, errs []errors.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errInfo := make([]ErrorInfo, len(errs))
	for i, err := range errs {
		info := ErrorInfo{
			Message: err.Message,
			Field:   err.Field,
		}
		errInfo[i] = info
	}

	json.NewEncoder(w).Encode(&Body{
		Errors: errInfo,
		Meta: Meta{
			HTTPStatus: status,
		},
	})
}

func Errors(status int, w http.ResponseWriter, errs []errors.Error) {
	SendError(status, w, errs)
}

func Error(w http.ResponseWriter, err errors.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)

	json.NewEncoder(w).Encode(&Body{
		Errors: []ErrorInfo{
			ErrorInfo{
				Message: err.Message,
			},
		},
		Meta: Meta{
			HTTPStatus: err.Code,
		},
	})
}
