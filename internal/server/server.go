package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"synapse-ai/internal/rag"
)

type Server struct {
	addr string
}

type LLMRequest struct {
	Prompt     string             `json:"prompt"`
	MsgHistory []*rag.ChatMessage `json:"msgHistory"`
}

type LLMResponse struct {
	Response   string             `json:"response"`
	MsgHistory []*rag.ChatMessage `json:"msgHistory"`
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

	userPrompt := req.Prompt
	history := req.MsgHistory

	res, history, err := rag.CallWithContext(userPrompt, history)

	if err != nil {
		log.Printf("LLM call failed: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	jsonResponse := LLMResponse{
		Response:   res,
		MsgHistory: history,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}

func (s Server) Run() error {
	log.Printf("Starting Server on %s", s.addr)
	return http.ListenAndServe(s.addr, s.routes())
}
