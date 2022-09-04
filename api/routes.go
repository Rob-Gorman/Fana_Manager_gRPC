package api

import "net/http"

func (s *Server) providerRoutes() {
	s.HandleFunc("/flagset", s.H.GetFlagset).Methods("GET")
}

func (s *Server) staticRoutes() {
	s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./build/static/")))).Methods("GET")
	s.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./build/index.html")
	}).Methods("GET")
}
