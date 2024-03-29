// Package random_test contains tests for the package random.
package random_test

import (
	"reflect"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/random"
)

// TestNewSeed tests the creation of a new seed using the function NewSeed.
func TestNewSeed(t *testing.T) {
	t.Parallel()
	if got := random.NewSeed(); reflect.DeepEqual(got, random.Seed{}) {
		t.Errorf("got: %v", got)
	}
}

// TestSeedFrom tests the creation of a seed from a string using the function SeedFrom.
func TestSeedFrom(t *testing.T) {
	t.Parallel()
	data := []struct {
		name   string
		source string
	}{
		{name: "source with letters is valid", source: "Test"},
		{name: "source with emoji is valid", source: "🐞"},
		{name: "source containing whitespace is valid", source: "Test/n🐞"},
		{name: "source with space is valid", source: " "},
		{name: "empty source is valid", source: ""},
	}
	for _, tt := range data {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := random.SeedFrom(tt.source); reflect.DeepEqual(got, random.Seed{}) {
				t.Errorf("got: %v", got)
			}
		})
	}
}
