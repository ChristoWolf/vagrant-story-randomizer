// Package psximager_tests the psximager package.
package psximager_test

import (
	"fmt"
	"reflect"
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

// ExampleExecute executes the psxinject command.
// Modify paths to match your environment.
func ExampleExecute() {
	err := psximager.Execute("C:\\Program Files\\psximager", "cueFile.cue", "orgFile", "injectFile")
	if err != nil {
		fmt.Println(err)
	}
	// Output: error executing psxinject: exit status 1
}
