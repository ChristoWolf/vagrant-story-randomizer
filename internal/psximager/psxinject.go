// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import "os/exec"

const (
	// PsxInject is the name of to the psxinject executable.
	PsxInject = "psxinject"
)

func NewInjectCmd(args ...string) *exec.Cmd {
	return NewPipedCmd(PsxInject, args...)
}
