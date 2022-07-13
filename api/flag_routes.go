package api

import (
	"sovereign/controllers"
)

func (s *Server) FlagRoutes() {
	s.HandleFunc("/flags", controllers.GetFlags(s.DB)).Methods("GET")
}
