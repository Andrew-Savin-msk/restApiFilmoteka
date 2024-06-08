package apiserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type server struct {
	mux    http.ServeMux
	store  string // Temporary
	logger logrus.Logger
}
