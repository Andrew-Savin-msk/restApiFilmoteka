package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	model "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	ctxUserKey = iota
	ctxRequestKey
)

var (
	sessionName = "user-key"
)

func (s *server) setMuxer() {
	s.mux.HandleFunc("/create", s.basePaths(s.handleCreateUser()))
	s.mux.HandleFunc("/get-session", s.basePaths(s.))
}

func (s *server) basePaths(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setting all middleware required for basic paths
		next = s.wrapSetRequestId(next)
		next = s.wrapLogRequest(next)
		next.ServeHTTP(w, r)
	})
}

func (s *server) protectedPaths(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next = s.wrapAuthorisation(next)
		next.ServeHTTP(w, r)
	})
}

func (s *server) wrapAuthorisation(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (s *server) wrapSetRequestId(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxRequestKey, id)))
	})
}

func (s *server) wrapLogRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxRequestKey),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)
		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %v %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}

func (s *server) handleCreateUser() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:  req.Email,
			Passwd: req.Password,
		}

		err = s.store.User().Create(u)
		if err != nil {
			s.errorResponse(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleGetSession() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		
	}

}

func (s *server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
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
