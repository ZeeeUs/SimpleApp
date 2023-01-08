package http

func InitRoutes(s *Server) {
	s.router.HandleFunc("/api/v1/list", s.BookHandler.GetList).Methods("GET")
	s.router.HandleFunc("/api/v1/book/{id:[0-9]+}", s.BookHandler.DeleteBook).Methods("DELETE", "OPTIONS")
	s.router.HandleFunc("/api/v1/book/add", s.BookHandler.AddBook).Methods("POST")
}
