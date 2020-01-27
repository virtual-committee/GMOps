package command

import (
	"fmt"

	"GMOps/src/shell/command/spec"
)

func buildGMOpsOriginalSSHCommand(args []string) (spec.OriginalSSHCommand, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command")
	}
	switch args[0] {
	case "create-repo":
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid command")
		}
		return &spec.CreateRepo{Name: args[1]}, nil
	}
	return &spec.MissedHit{}, nil
}

func BuildOriginalSSHCommand(args []string) (spec.OriginalSSHCommand, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("invalid command")
	}
	switch args[0] {
	case "git-receive-pack":
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid command")
		}
		return spec.NewGitReceivePack(args[1]), nil
	case "git-upload-pack":
		if len(args) != 2 {
			return nil, fmt.Errorf("invalid command")
		}
		return spec.NewGitUploadPack(args[1]), nil
	case "gmops":
		return buildGMOpsOriginalSSHCommand(args[1:])
	}

	return &spec.MissedHit{}, nil
}
