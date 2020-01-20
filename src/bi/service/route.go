package service

func (s *Service) initRoute() error {

	userRouter := s.r.Group("/user")
	{
		userRouter.POST("/register", s.createUserAction)
	}

	return nil
}
