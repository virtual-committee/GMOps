package app

import (
	"context"

	"GMOps/src/app/options"
	bi "GMOps/src/bi/service"
)

func Run(ctx context.Context, cancel context.CancelFunc, opt *options.ServerOption) error {
	biServer := bi.NewService(ctx, cancel, opt.BIAddrPath, opt.MongoConnector)
	biServer.Run()

	select {
	case <-ctx.Done():
	}
	return nil
}
