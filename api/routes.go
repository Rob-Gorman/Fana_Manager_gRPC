package api

import (
	"embed"
	"io/fs"
	"net/http"
)

func (s *Server) providerRoutes() {
	s.HandleFunc("/flagset", s.H.GetFlagset).Methods("GET")
}

func (s *Server) staticRoutes(dash embed.FS, index []byte) {
	staticSubDir, _ := fs.Sub(dash, "static")
	staticFS := http.FS(staticSubDir)
	s.Handle("/", http.FileServer(staticFS)).Methods("GET")
	s.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(staticFS))).Methods("GET")
	s.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(index)
	}).Methods("GET")
}
