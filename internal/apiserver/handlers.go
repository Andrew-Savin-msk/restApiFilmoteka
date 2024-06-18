package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	actor "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/actor"
	user "github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/model/user"
)

const (
	ctxUserKey = iota
	ctxRequestKey
)

var (
	sessionName = "user-key"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not auntificated")
)

func (s *server) handleCreateUser() http.HandlerFunc {

	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.errorResponse(w, r, http.StatusMethodNotAllowed, nil)
			return
		}

		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		u := &user.User{
			Email:   req.Email,
			Passwd:  req.Password,
			IsAdmin: false,
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
		Email  string `json:"email"`
		Passwd string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.errorResponse(w, r, http.StatusMethodNotAllowed, nil)
			return
		}

		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		err = u.CompareHashAndPassword(req.Passwd)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.errorResponse(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.Id
		err = session.Save(r, w)
		if err != nil {
			s.errorResponse(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}

}

func (s *server) handleWhoamI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			s.errorResponse(w, r, http.StatusMethodNotAllowed, nil)
			return
		}

		u, ok := r.Context().Value(ctxUserKey).(*user.User)
		if !ok {
			s.errorResponse(w, r, http.StatusUnprocessableEntity, nil)
			return
		}

		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) handleCreateActor() http.HandlerFunc {
	type request struct {
		Name      string `json:"name"`
		Gen       string `json:"gender"`
		Birthdate string `json:"birthdate"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			s.errorResponse(w, r, http.StatusMethodNotAllowed, nil)
			return
		}

		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		birth, err := time.Parse("01-02-2006", req.Birthdate)
		if err != nil {
			s.errorResponse(w, r, http.StatusBadRequest, err)
			return
		}

		act := &actor.Actor{
			Name:      req.Name,
			Gen:       req.Gen,
			Birthdate: birth,
		}
		s.logger.Infof("Time: %v", act)

		s.respond(w, r, http.StatusOK, nil)
	}
}

// Interface methods

func (s *server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
