// Package psximager_tests the psximager package.
package psximager_test

import (
	"reflect"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/psximager"
)

// TestNewInjectCmd tests the NewInjectCmd function.
func TestNewInjectCmd(t *testing.T) {
	t.Parallel()
	wantCmd := psximager.PsxInject
	wantArgs := []string{"cueFile.cue", "orgFile", "injectFile"}
	want := psximager.NewPipedCmd(wantCmd, wantArgs...)
	// Test that the command matches our expectations.
	got := psximager.NewInjectCmd(wantArgs[0], wantArgs[1], wantArgs[2])
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("got: %v, want: %v", *got, *want)
	}
}
