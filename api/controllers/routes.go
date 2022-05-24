package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", s.UpdateUser)
		v1.PUT("/avatar/users/:id", s.UpdateAvatar)
		v1.DELETE("/users/:id", s.DeleteUser)

	}
}
