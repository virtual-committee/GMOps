package app

import (
	"context"

	"github.com/gin-gonic/gin"

	"GMOps/src/app/options"
	bi "GMOps/src/bi/service"
)

func Run(ctx context.Context, cancel context.CancelFunc, opt *options.ServerOption) error {
	if !opt.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	biServer, err := bi.NewService(ctx, cancel, opt.BIAddrPath, opt.MongoConnector)
	if err != nil {
		return err
	}
	biServer.Run()

	select {
	case <-ctx.Done():
	}
	return nil
}
