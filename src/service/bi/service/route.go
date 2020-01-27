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
			authorizedUserRouter.GET("/repo/:name", s.getUserRepoAction)
			authorizedUserRouter.POST("/repo", s.createUserRepoAction)
		}
	}

	keyRouter := s.r.Group("/key")
	{
		keyRouter.POST("/:id/apply", s.applyUserAuthKeyAction)
		keyRouter.POST("/:id/cancel", s.cancelUserAuthKeyAction)
	}

	repoRouter := s.r.Group("/repo")
	{
		repoRouter.GET("/:repoId/hook/pre-receive", s.getRepoPreReceiveHooksAction)
		repoRouter.POST("/:repoId/hook/:hookId/apply", s.useRepoHookAction)
	}

	hookRouter := s.r.Group("/hook")
	{
		hookRouter.POST("/add", s.addGitHookAction)
	}

	return nil
}
