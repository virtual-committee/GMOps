package spec

import (
	"io"
	"os"
	"os/exec"
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

func (g *GitReceivePack) SetUser(user string) { g.user = user }

func (g *GitReceivePack) SetReadWriter(stdin io.Reader, stdout, stderr io.Writer) {
	g.stdin = stdin
	g.stdout = stdout
	g.stderr = stderr
}

func (g *GitReceivePack) Exec() error {
	if err := g.repo.Parse(); err != nil {
		return err
	}

	cmd := exec.Command("git", "receive-pack", GMOPS_REPO_DATA_PATH+g.repo.id)
	cmd.Stdin = g.stdin
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr
	cmd.Env = append(os.Environ(), "GMOPS_USER="+g.user)
	cmd.Env = append(cmd.Env, "GMOPS_REPO_ID="+g.repo.id)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
