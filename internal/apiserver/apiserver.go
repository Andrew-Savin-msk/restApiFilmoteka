package apiserver

import (
	"net/http"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/config"
)

func New(cfg config.Config) *server {
	return &server{
		mux: newMux(),
		logger: 
	}
}

func newMux() http.ServeMux {

}

func setLog(level string) {
	
}