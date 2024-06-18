package apiserver

import (
	"net/http"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/config"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/store"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

type server struct {
	mux          *http.ServeMux
	sessionStore sessions.Store
	store        store.Store
	logger       *logrus.Logger
}

func newServer(st store.Store, cfg *config.Config) *server {
	srv := &server{
		mux:          http.NewServeMux(),
		logger:       setLog(cfg.LogLevel),
		store:        st,
		sessionStore: sessions.NewCookieStore([]byte(cfg.SessionKey)),
	}

	srv.setMuxer()

	return srv
}

func (s *server) setMuxer() {
	s.mux.Handle("/register", s.basePaths(s.handleCreateUser()))
	s.mux.Handle("/authorize", s.basePaths(s.handleGetSession()))
	s.mux.Handle("/who-am-i", s.protectedPaths(s.handleWhoamI()))
	s.mux.Handle("/private/create-actor", s.protectedPaths(s.handleCreateActor()))
}
