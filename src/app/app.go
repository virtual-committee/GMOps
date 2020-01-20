package app

import (
	"context"

	"github.com/gin-gonic/gin"

	"GMOps/src/app/options"
	bi "GMOps/src/bi/service"
)

func Run(opt *options.ServerOption) error {
	if !opt.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	ctx, cancel := context.WithCancel(context.Background())

	biServer, err := bi.NewService(ctx, cancel, opt.BIAddrPath, opt.MongoConnector, opt.DBName)
	if err != nil {
		return err
	}

	if opt.InitDB {
		if err = biServer.InitDB(); err != nil {
			return err
		}
	} else {
		if err = biServer.Run(); err != nil {
			return err
		}

		select {
		case <-ctx.Done():
		}
	}

	return nil
}
