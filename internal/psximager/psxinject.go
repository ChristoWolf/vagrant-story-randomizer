// Code contained in the psxinject.go file
// wraps the psxinject tool which
// is part of the psximager tool suite.
package psximager

import (
	"fmt"
	"os/exec"
)

const (
	// PsxInject is the name of to the psxinject executable.
	PsxInject = "psxinject"
)

// newInjectCmd returns a new exec.Cmd that is configured to
// run the psxinject in a piped manner.
func newInjectCmd(psximagerPath, cueFile, orgFile, injectFile string) *exec.Cmd {
	return NewPipedCmd(psximagerPath, PsxInject, cueFile, orgFile, injectFile)
}

// Execute runs an inject command
// given a cue file, the original file to overwrite inside the image,
// and the file which shall be injected.
func Execute(psximagerPath, cueFile, orgFile, injectFile string) error {
	cmd := newInjectCmd(psximagerPath, cueFile, orgFile, injectFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing %s: %w", PsxInject, err)
	}
	return nil
}
