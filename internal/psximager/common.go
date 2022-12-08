// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import (
	"os"
	"os/exec"
)

// NewPipedCmd returns a new exec.Cmd that is configured to
// run a named command in a piped manner
// inside the given working directory.
func NewPipedCmd(name string, args ...string) *exec.Cmd {
	prefix := []string{"start", "/min", "/c"}
	fullArgs := append(prefix, name)
	fullArgs = append(fullArgs, args...)
	cmd := exec.Command("cmd.exe", fullArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
