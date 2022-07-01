// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import "os/exec"

const (
	// PsxRip is the name of to the psxrip executable.
	PsxRip = "psxrip"
)

func NewRipCmd(args ...string) *exec.Cmd {
	return NewPipedCmd(PsxRip, args...)
}
