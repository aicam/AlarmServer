package server

func (s *Server) Routes() {
	s.Router.POST("/add_info", s.AddInfo())
	s.Router.GET("/get_info", s.GetInfo())
}
