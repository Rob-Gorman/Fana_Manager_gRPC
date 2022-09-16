package api

import (
	"embed"
	"manager/database"
	"manager/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/veqryn/h2c"
	"golang.org/x/net/http2"
)

type Server struct {
	*mux.Router
	H handlers.Handler
	// this Handler is a wrapper around our DB to allow us to define methods on it
	// those methods being our controller functions (handlers)
}

func NewServer(dash embed.FS, index []byte) *Server {
	s := &Server{
		H:      handlers.New(database.Init()),
		Router: mux.NewRouter(),
	}

	s.routes(dash, index)
	return s
}

func NewH2Server(r *mux.Router, port string) http.Server {
	h2Handle := &h2c.HandlerH2C{
		Handler:  r,
		H2Server: &http2.Server{},
	}

	return http.Server{
		Addr:    port,
		Handler: h2Handle,
	}
}

func (s *Server) routes(dash embed.FS, index []byte) {
	s.providerRoutes()
	s.staticRoutes(dash, index)
}
