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
	// Public endpoints
	s.mux.Handle("/register", s.basePaths(s.handleCreateUser()))
	s.mux.Handle("/authorize", s.basePaths(s.handleGetSession()))
	s.mux.Handle("/get-actor", s.basePaths(s.handleGetActor()))
	s.mux.Handle("/get-actors", s.basePaths(s.handleGetActors()))

	// Authorisation required endpoints
	s.mux.Handle("/private/who-am-i", s.protectedPaths(s.handleWhoamI()))

	// Admin rights required endpoints
	s.mux.Handle("/private/create-actor", s.adminPaths(s.handleCreateActor()))
	s.mux.Handle("/private/delete-actor", s.adminPaths(s.handleDeleteActor()))
	s.mux.Handle("/private/update-actor", s.adminPaths(s.handleOverwrightActor()))
	s.mux.Handle("/private/post-film", s.adminPaths(s.handleCreateFilm()))
}
