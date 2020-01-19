package main

import (
	"context"
	"os"
	"runtime"

	"github.com/spf13/pflag"

	"GMOps/src/app"
	"GMOps/src/app/options"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	opt := options.NewServerOption()
	opt.Add(pflag.CommandLine)

	ctx, cancel := context.WithCancel(context.Background())

	if err := app.Run(ctx, cancel, opt); err != nil {
		os.Exit(1)
	}
}
