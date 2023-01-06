package http

func InitRoutes(s *Server) {
	s.router.HandleFunc("/api/v1/list", s.BookHandler.BookHandler).Methods("GET")
}
