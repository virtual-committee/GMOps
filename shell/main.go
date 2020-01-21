package main

import (
	"os"
	"runtime"

	"GMOps/shell/command"

	"github.com/mattn/go-shellwords"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser := shellwords.NewParser()
	line := os.Getenv("SSH_ORIGINAL_COMMAND")
	args, err := parser.Parse(line)
	if err != nil {
		os.Exit(1)
	}

	originalCommand, err := command.BuildOriginalSSHCommand(args)
	if err != nil {
		os.Exit(1)
	}
	if originalCommand.RequiredUser() {
		principal, err := command.ParsePrincipal(os.Args[1:])
		if err != nil {
			os.Exit(1)
		}
		valid, err := principal.Valid()
		if err != nil {
			os.Exit(1)
		}
		if !valid {
			os.Exit(1)
		}
		originalCommand.SetUser(principal.Username)
	}

	originalCommand.SetReadWriter(os.Stdin, os.Stdout, os.Stderr)

	if err := originalCommand.Exec(); err != nil {
		os.Exit(1)
	}
}
