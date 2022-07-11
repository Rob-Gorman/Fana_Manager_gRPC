package api

import (
	"encoding/json"
	"net/http"
	"sovereign/data"

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
	s.HandleFunc("/", s.getMongoData()).Methods("GET")
}

func (s *Server) getMongoData() http.HandlerFunc {
	query := data.ConnectDB()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(query); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
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
