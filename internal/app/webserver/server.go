package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/middlewares"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/store"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/utils"
)

type server struct {
	router     *mux.Router
	logger     *logrus.Logger
	store      store.Store
	redisStore store.RedisStore
}

func newServer(store store.Store, rdb store.RedisStore) *server {
	s := &server{
		router:     mux.NewRouter(),
		logger:     logrus.New(),
		store:      store,
		redisStore: rdb,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/registration", s.handleUsersRegistration()).Methods("POST")
	s.router.HandleFunc("/auth", s.handleUsersAuth()).Methods("POST")
	s.router.HandleFunc("/time", middlewares.Authenticate(s.handleTime())).Methods("GET")
	s.router.HandleFunc("/refresh-token", s.handleRefreshToken()).Methods("GET")
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, _ *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) handleUsersRegistration() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
			Username: req.Username,
		}

		err = s.store.User().Create(u)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleUsersAuth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type respond struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, store.ErrIncorrectEmailOrPassword)
			return
		}

		aToken, rToken, err := utils.GenerateJwt(u)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.redisStore.Set(aToken, rToken)

		res := &respond{
			Token: fmt.Sprintf("Bearer %s", aToken),
		}

		s.respond(w, r, http.StatusOK, res)
	}
}

func (s *server) handleRefreshToken() http.HandlerFunc {
	type respond struct {
		Token string `json:"token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.ExtractTokenFromHandler(r)
		_, err := s.redisStore.Get(tokenString)
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, err)
			return
		}

		if err := s.redisStore.Delete(tokenString); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		aToken, rToken, err := utils.RefreshJwt(tokenString)

		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.redisStore.Set(aToken, rToken)

		res := &respond{
			Token: fmt.Sprintf("Bearer %s", aToken),
		}

		s.respond(w, r, http.StatusOK, res)
	}
}

func (s *server) handleTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC1123)

		s.respond(w, r, http.StatusOK, map[string]string{"time": currentTime})
	}
}
