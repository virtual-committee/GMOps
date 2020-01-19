package service

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Service struct {
	r              *gin.Engine
	unixSocketPath string
	mongoConnector string
	ctx            context.Context
	cancel         context.CancelFunc
}

func NewService(ctx context.Context, cancel context.CancelFunc, unixSocketPath string, mongoConnector string) *Service {
	s := &Service{
		r:              gin.Default(),
		unixSocketPath: unixSocketPath,
		mongoConnector: mongoConnector,
		ctx:            ctx,
		cancel:         cancel,
	}

	return s
}

func (s *Service) Run() error {
	go func() {
		s.r.RunUnix(s.unixSocketPath)
		s.cancel()
	}()
	return nil
}
