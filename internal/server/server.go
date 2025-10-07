package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger 	  *log.Logger
	HttpServer http.Server
}

func NewServer(log *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.GetIndex)
	mux.HandleFunc("/upload", handlers.PostForm)

	s := http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: log,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	return &Server{
		Logger: log,
		HttpServer: s,
	}
}