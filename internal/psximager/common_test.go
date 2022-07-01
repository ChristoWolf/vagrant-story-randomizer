// Package psximager_tests the psximager package.
package psximager_test

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/psximager"
)

// TestNewPipedCmd tests the NewPipedCmd function.
func TestNewPipedCmd(t *testing.T) {
	t.Parallel()
	wantCmd := "test-cmd"
	wantArgs := []string{"-test", "args"}
	want := exec.Command(wantCmd, wantArgs...)
	want.Stdout = os.Stdout
	want.Stderr = os.Stderr
	// Test that the command matches our expectations.
	got := psximager.NewPipedCmd("test-cmd", wantArgs...)
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("got: %v, want: %v", *got, *want)
	}
}
