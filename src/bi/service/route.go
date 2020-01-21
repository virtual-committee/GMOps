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
			authorizedUserRouter.GET("/repos", s.getUserReposAction)
			authorizedUserRouter.POST("/repo", s.createUserRepoAction)
		}
	}

	keyRouter := s.r.Group("/key")
	{
		keyRouter.POST("/:id/apply", s.applyUserAuthKeyAction)
		keyRouter.POST("/:id/cancel", s.cancelUserAuthKeyAction)
	}

	return nil
}
