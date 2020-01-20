package service

func (s *Service) initRoute() error {

	userRouter := s.r.Group("/user")
	{
		userRouter.POST("/register", s.createUserAction)
		authorizedUserRouter := userRouter.Use(s.authorize())
		{
			authorizedUserRouter.GET("/valid", s.validUserAction)
		}
	}

	return nil
}
