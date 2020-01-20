package service

func (s *Service) initRoute() error {

	userRouter := s.r.Group("/user")
	{
		userRouter.POST("/add", s.createUserAction)
		authorizedUserRouter := userRouter.Use(s.authorize())
		{
			authorizedUserRouter.GET("/valid", s.validUserAction)
			authorizedUserRouter.GET("/keys", s.getUserAuthKeysAction)
			authorizedUserRouter.POST("/key", s.addUserAuthKeyAction)
			authorizedUserRouter.GET("/key/:id/valid", s.validUserAuthKeyAction)
		}
	}

	return nil
}
