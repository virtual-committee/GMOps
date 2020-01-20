package main

import (
	"flag"
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

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := app.Run(opt); err != nil {
		os.Exit(1)
	}
}
