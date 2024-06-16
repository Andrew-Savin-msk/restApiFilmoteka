package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/models/user"
)

func (s *server) setMuxer() {
	s.mux.HandleFunc("/create", s.handleCreateUser())
}

func (s *server) handleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &user.Request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
		}

		// TODO: Db inter

		s.respond(w, r, http.StatusOK, "")
	}
}

func (s *server) errorResponse(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
