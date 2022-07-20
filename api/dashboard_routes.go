package api

import "github.com/gorilla/mux"

func (s *Server) dashboardRoutes() *mux.Route {
	s.HandleFunc("/api/flags", s.H.GetAllFlags).Methods("GET")
	s.HandleFunc("/api/flags", s.H.CreateFlag).Methods("POST")
	s.HandleFunc("/api/flags/{id}", s.H.GetFlag).Methods("GET")
	// s.HandleFunc("/flags/{id}", s.H.UpdateFlag).Methods("PATCH")
	s.HandleFunc("/api/audiences", s.H.GetAllAudiences).Methods("GET")
	s.HandleFunc("/api/audiences/{id}", s.H.GetAudience).Methods("GET")
	s.HandleFunc("/api/attributes", s.H.GetAllAttributes).Methods("GET")
	s.HandleFunc("/api/attributes", s.H.CreateAttribute).Methods("POST")
}
