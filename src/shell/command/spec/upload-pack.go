package spec

import (
	"io"
	"os/exec"
)

type GitUploadPack struct {
	repo   Repo
	user   string
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

var _ OriginalSSHCommand = &GitUploadPack{}

func NewGitUploadPack(originalRepoName string) *GitUploadPack {
	return &GitUploadPack{repo: Repo{OriginalRepoName: originalRepoName}}
}

func (g *GitUploadPack) RequiredUser() bool { return true }

func (g *GitUploadPack) SetUser(user string) { g.user = user }

func (g *GitUploadPack) SetReadWriter(stdin io.Reader, stdout, stderr io.Writer) {
	g.stdin = stdin
	g.stdout = stdout
	g.stderr = stderr
}

func (g *GitUploadPack) Exec() error {
	if err := g.repo.Parse(); err != nil {
		return err
	}

	cmd := exec.Command("git", "upload-pack", GMOPS_REPO_DATA_PATH+g.repo.id)
	cmd.Stdin = g.stdin
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
