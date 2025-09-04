package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.homeHandler)
	r.Get("/healthCheck", s.healthHandler)

	return r
}

func CreateServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok"}`)
}

func (s Server) Run() error {
	log.Printf("Starting Server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.routes())
}
