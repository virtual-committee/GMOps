package main

import (
	"flag"
	"os"
	"runtime"

	"GMOps/src/service/app"
	"GMOps/src/service/app/options"

	"github.com/spf13/pflag"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	opt := options.NewServerOption()
	opt.Add(pflag.CommandLine)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := app.Run(opt); err != nil {
		os.Exit(1)
	}
}
