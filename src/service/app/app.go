package app

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"

	"GMOps/src/service/app/options"
	bi "GMOps/src/service/bi/service"
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
		os.SetEnv("GMOPS_BI_UNIX_SOCKET", opt.BIAddrPath)

		select {
		case <-ctx.Done():
		}
	}

	return nil
}
