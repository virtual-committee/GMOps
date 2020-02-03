package service

func (s *Service) initRoute() error {

	userRouter := s.r.Group("/user")
	{
		userRouter.POST("/add", s.createUserAction)
		authorizedUserRouter := userRouter.Use(s.authorizeMiddleware())
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

	repoRouter := s.r.Group("/repo/:repoId").Use(s.repoMiddleware())
	{
		repoRouter.GET("/hooks/:hookType", s.getRepoHooksAction)
		repoRouter.POST("/hook/:hookId/apply", s.useRepoHookAction)
		repoRouter.GET("/refs", s.getUserRepoRefsAction)
	}
	repoCommitRouter := s.r.Group("/repo/:repoId/commit/:commitId").Use(s.repoMiddleware()).Use(s.repoCommitMiddleware())
	{
		repoCommitRouter.GET("/info", s.getCommitInfoAction)
	}

	hookRouter := s.r.Group("/hook")
	{
		hookRouter.POST("/add", s.addGitHookAction)
	}

	return nil
}
