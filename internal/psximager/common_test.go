// Package psximager_test tests the psximager package.
package psximager_test

import (
	"os"
	"os/exec"
	"path"
	"reflect"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/psximager"
)

// TestNewPipedCmd tests the NewPipedCmd function.
func TestNewPipedCmd(t *testing.T) {
	t.Parallel()
	wantPrefix := []string{"start", "/min", "/c"}
	wantPwd := "test/path"
	wantCmd := "test-cmd"
	wantArgs := []string{"-test", "args"}
	wantFullArgs := append(wantPrefix, path.Join(wantPwd, wantCmd))
	wantFullArgs = append(wantFullArgs, wantArgs...)
	want := exec.Command("cmd.exe", wantFullArgs...)
	want.Stdout = os.Stdout
	want.Stderr = os.Stderr
	// Test that the command matches our expectations.
	got := psximager.NewPipedCmd(wantPwd, wantCmd, wantArgs...)
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("got: %v, want: %v", *got, *want)
	}
}
