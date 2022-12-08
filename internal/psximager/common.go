// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import (
	"bytes"
	"io"
	"os/exec"
	"path"
)

// NewPipedCmd returns a new exec.Cmd that is configured to
// run a named command in a piped manner
// inside the given working directory.
func NewPipedCmd(pwd string, name string, args ...string) *exec.Cmd {
	prefix := []string{"start", "/min", "/c"}
	fullArgs := append(prefix, path.Join(pwd, name))
	fullArgs = append(fullArgs, args...)
	cmd := exec.Command("cmd.exe", fullArgs...)
	// For now, we don't need stdout.
	// Might change in the future.
	cmd.Stdout = io.Discard
	// Buffer stderr s.t. it can be read
	// in case of an error.
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	return cmd
}
