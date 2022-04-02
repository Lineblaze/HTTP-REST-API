package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//APIServer ...

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//Start ...

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return nil
	}

	s.configureRouter()

	s.logger.Info("starting api server")
	
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Logger

func (s *APIServer) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// Router

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, "Hello")
	}
}