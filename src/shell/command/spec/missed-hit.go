package spec

import (
	"fmt"
	"io"
)

type MissedHit struct{}

var _ OriginalSSHCommand = &MissedHit{}

func (c *MissedHit) Exec() error {
	return fmt.Errorf("missed hit")
}

func (c *MissedHit) SetReadWriter(_ io.Reader, _, _ io.Writer) {}

func (c *MissedHit) RequiredUser() bool { return false }

func (c *MissedHit) SetUser(_ string) {}
