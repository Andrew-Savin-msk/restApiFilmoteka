package apiserver

import (
	"net/http"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	mux    *http.ServeMux
	store  store.Store // Temporary
	logger *logrus.Logger
}
