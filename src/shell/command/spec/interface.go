package spec

import (
	"io"
)

type OriginalSSHCommand interface {
	Exec() error
	SetReadWriter(stdIn io.Reader, stdOut, stdErr io.Writer)
	RequiredUser() bool
	SetUser(user string)
}
