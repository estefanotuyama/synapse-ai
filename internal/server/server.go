package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
	"synapse-ai/internal/rag"
)

type Server struct {
	addr string
}

type LLMRequest struct {
	Prompt     string             `json:"prompt"`
	MsgHistory []*rag.ChatMessage `json:"MsgHistory"`
}

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.homeHandler)
	r.Get("/healthCheck", s.healthHandler)
	r.Post("/llm_call", s.llmHandler)

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

func (s *Server) llmHandler(w http.ResponseWriter, r *http.Request) {
	var req LLMRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Println(req.Prompt)

	res, _, err := rag.CallWithContext(req.Prompt, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Type of res: %v", reflect.TypeOf(res)) //*genai.GenerateContentResponse
}

func (s Server) Run() error {
	log.Printf("Starting Server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.routes())
}
