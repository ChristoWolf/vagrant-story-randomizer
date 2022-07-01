// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import "os/exec"

const (
	// PsxBuild is the name of to the psxbuild executable.
	PsxBuild = "psxbuild"
)

func NewBuildCmd(args ...string) *exec.Cmd {
	return NewPipedCmd(PsxBuild, args...)
}
