package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/anatolethien/forum/internal/app/repository"
	"github.com/anatolethien/forum/internal/app/service"

	"github.com/anatolethien/forum/internal/app/handler"
)

type Server struct {
	httpServer *http.Server
}

func New(config *Config) *Server {
	db, err := repository.OpenDB(config.DBDriver, config.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.Addr
	}

	return &Server{
		httpServer: &http.Server{
			Addr:         port,
			Handler:      handler.InitRouter(),
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("starting api server at", s.httpServer.Addr)

	return s.httpServer.ListenAndServe()
}
