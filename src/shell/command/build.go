package command

import (
	"fmt"

	"GMOps/shell/command/spec"
)

func buildGitOriginalSSHCommand(args []string) (spec.OriginalSSHCommand, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command")
	}
	switch args[0] {
	case "receive-pack":
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid command")
		}
		return spec.NewGitReceivePack(args[1]), nil
	case "upload-pack":
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid command")
		}
		return spec.NewGitUploadPack(args[1]), nil
	}
	return &spec.MissedHit{}, nil
}

func BuildOriginalSSHCommand(args []string) (spec.OriginalSSHCommand, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command")
	}
	switch args[0] {
	case "git":
		return buildGitOriginalSSHCommand(args[1:])
	}

	return &spec.MissedHit{}, nil
}
