package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Server struct {
	host        string
	server      *http.Server
	router      *mux.Router
	BookHandler BookHandler
}

type BookHandler interface {
	GetList(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
}

func (s *Server) WithBookHandler(handler BookHandler) *Server {
	s.BookHandler = handler
	return s
}

func (s *Server) Run(corsSettings *cors.Cors) (err error) {
	InitRoutes(s)

	server := &http.Server{
		Addr:         s.host,
		Handler:      corsSettings.Handler(s.router),
		WriteTimeout: http.DefaultClient.Timeout,
		ReadTimeout:  http.DefaultClient.Timeout,
	}

	s.server = server
	go func() {
		if err = server.ListenAndServe(); err != nil {
			return
		}
	}()
	return nil
}

func (s *Server) Shutdown() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err = s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return
}

func NewServer(host string) *Server {
	return &Server{
		host:   host,
		router: mux.NewRouter(),
	}
}
