package api

import (
	"manager/database"
	"manager/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router // this is our express-router
	H           handlers.Handler
	// this Handler is a wrapper around our DB to allow us to define methods on it
	// those methods being our controller functions (handlers)
}

func NewServer() *Server {
	s := &Server{
		H:      handlers.New(database.Init()),
		Router: mux.NewRouter(),
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.dashboardRoutes()
	// s.providerRoutes()
}
