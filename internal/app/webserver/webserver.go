package webserver

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// WebServer ...
type WebServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *WebServer {
	return &WebServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *WebServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting web server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *WebServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *WebServer) configureRouter() {
	s.router.HandleFunc("/time", s.handleTime())
}

func (s *WebServer) handleTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, time.Now().Format(time.RFC1123))
	}
}
