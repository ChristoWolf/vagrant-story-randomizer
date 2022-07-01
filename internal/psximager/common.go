// Package psximager provides a simple API wrapper
// around psximager's psxbuild tool.
// See https://github.com/cebix/psximager.
// This package is not intended to be used directly.
package psximager

import (
	"os"
	"os/exec"
)

func NewPipedCmd(name string, args ...string) *exec.Cmd {
	prefix := []string{"start", "/min", "/c"}
	fullArgs := append(prefix, name)
	fullArgs = append(fullArgs, args...)
	cmd := exec.Command("cmd.exe", fullArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
