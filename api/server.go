package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer() *Server {
	s := &Server{
		Ruleset: &Rules,
		Router:  mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/", s.getFlagData()).Methods("GET")
}

func (s *Server) getFlagData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(*(s.Ruleset)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
