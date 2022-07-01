// Package psximager_tests the psximager package.
package psximager_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/psximager"
)

// TestNewRipCmd tests the NewRipCmd function.
func TestNewRipCmd(t *testing.T) {
	t.Parallel()
	wantCmd := psximager.PsxRip
	wantArgs := []string{"-test", "args"}
	want := psximager.NewPipedCmd(wantCmd, wantArgs...)
	want.Stdout = os.Stdout
	want.Stderr = os.Stderr
	// Test that the command matches our expectations.
	got := psximager.NewRipCmd(wantArgs...)
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("got: %v, want: %v", *got, *want)
	}
}
