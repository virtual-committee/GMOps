package spec

import (
	"fmt"
	"io"
	"os/exec"
)

const (
	GIT_REPO_OWNER_I = 0o100
	GIT_REPO_OWNER_O = 0o200
	GIT_REPO_OWNER_X = 0o400
	GIT_REPO_GROUP_I = 0o010
	GIT_REPO_GROUP_O = 0o020
	GIT_REPO_GROUP_X = 0o040
	GIT_REPO_OTHER_I = 0o001
	GIT_REPO_OTHER_O = 0o002
	GIT_REPO_OTHER_X = 0x004
)

type GitReceivePack struct {
	repo   Repo
	user   string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

var _ OriginalSSHCommand = &GitReceivePack{}

func NewGitReceivePack(originalRepoName string) *GitReceivePack {
	return &GitReceivePack{repo: Repo{OriginalRepoName: originalRepoName}}
}

func (g *GitReceivePack) RequiredUser() bool { return true }

func (g *GitReceivePack) SetUser(user string) {
	g.user = user
}

func (g *GitReceivePack) SetReadWriter(stdin io.Reader, stdout, stderr io.Writer) {
	g.stdin = stdin
	g.stdout = stdout
	g.stderr = stderr
}

func (g *GitReceivePack) Exec() error {
	if err := g.repo.Parse(); err != nil {
		return err
	}
	if g.user == g.repo.user && g.repo.attr&GIT_REPO_OWNER_I == 0 {
		return fmt.Errorf("permission denied")
	} else if g.repo.attr&GIT_REPO_OTHER_I == 0 {
		return fmt.Errorf("permission denied")
	}

	cmd := exec.Command("git", "receive-pack", GMOPS_REPO_DATA_PATH+g.repo.id)
	cmd.Stdin = g.stdin
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(cmd.Args)
		return err
	}

	return nil
}
