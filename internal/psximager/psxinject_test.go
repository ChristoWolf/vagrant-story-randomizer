package psximager_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/psximager"
)

// TestNewInjectCmd tests the newInjectCmd function.
func TestNewInjectCmd(t *testing.T) {
	t.Parallel()
	wantPwd := "test/path"
	wantCmd := psximager.PsxInject
	wantArgs := []string{"cueFile.cue", "orgFile", "injectFile"}
	want := psximager.NewPipedCmd(wantPwd, wantCmd, wantArgs...)
	// Test that the command matches our expectations.
	got := psximager.NewInjectCmd(wantPwd, wantArgs[0], wantArgs[1], wantArgs[2])
	if !reflect.DeepEqual(*got, *want) {
		t.Errorf("got: %v, want: %v", *got, *want)
	}
}

// TestExecuteError tests the error handling
// behavior of the Execute function
// in case the command execution failed.
func TestExecuteError(t *testing.T) {
	t.Parallel()
	wrongPath := "./non-existent-path"
	err := psximager.Execute(wrongPath, "cueFile.cue", "orgFile", "injectFile")
	if err == nil {
		t.Errorf("expected a non-nil error")
	}
	got := err.Error()
	want := fmt.Sprintf("error executing %s: exit status 1: '%s' is not recognized", psximager.PsxInject, wrongPath)
	if strings.HasPrefix(got, want) {
		t.Errorf("got: %v, want substring: %v", got, want)
	}
}

// ExampleExecute executes the psxinject command.
// Modify paths to match your environment.
func ExampleExecute() {
	err := psximager.Execute("./psximager", "cueFile.cue", "orgFile", "injectFile")
	if err != nil {
		fmt.Println(err)
	}
}
