package api

func (s *Server) dashboardRoutes() {
	s.HandleFunc("/flags", s.H.GetAllFlags).Methods("GET")
	s.HandleFunc("/flags", s.H.CreateFlag).Methods("POST")
	s.HandleFunc("/flags/{id}", s.H.GetFlag).Methods("GET")
	s.HandleFunc("/audiences", s.H.GetAllAudiences).Methods("GET")
	s.HandleFunc("/audiences/{id}", s.H.GetAudience).Methods("GET")
	s.HandleFunc("/attributes", s.H.GetAllAttributes).Methods("GET")
	s.HandleFunc("/attributes", s.H.CreateAttribute).Methods("POST")
}
