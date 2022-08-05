package internal_test

import (
	"fmt"
	"github.com/christowolf/vagrant-story-randomizer/internal"
	"math"
	"testing"
)

func TestNewSeedRandomInt(t *testing.T) {
	t.Parallel()
	data := []struct {
		min     int
		max     int
		wantErr bool
	}{
		{min: 0, max: 69, wantErr: false},
		{min: 0, max: 0, wantErr: false},
		{min: 69, max: 0, wantErr: true},
		{min: 0, max: -69, wantErr: true},
		{min: -69, max: 0, wantErr: false},
		{min: 69, max: 69, wantErr: false},
		{min: 69, max: -69, wantErr: true},
		{min: -69, max: 69, wantErr: false},
		{min: -69, max: -69, wantErr: false},
		{min: math.MinInt, max: math.MaxInt, wantErr: false},
		{min: math.MaxInt, max: math.MinInt, wantErr: true},
	}
	for _, tt := range data {
		tt := tt
		t.Run(fmt.Sprintf("min: %d, max: %d", tt.min, tt.max), func(t *testing.T) {
			t.Parallel()
			seed := internal.NewSeed()
			got, err := seed.RandomInt(tt.min, tt.max)
			if err != nil && !tt.wantErr {
				t.Errorf("expected no error, got: %v", err)
				return
			}
			if !tt.wantErr && (got < tt.min || got > tt.max) {
				t.Errorf("got: %d; want: min: %d, max: %d", got, tt.min, tt.max)
			}
		})
	}
}

func TestSeedFromRandomInt(t *testing.T) {
	t.Parallel()
	data := []struct {
		source  string
		min     int
		max     int
		wantErr bool
	}{
		{source: "Test", min: 0, max: 69, wantErr: false},
		{source: "Test", min: 69, max: 0, wantErr: true},
		{source: "🐞", min: 0, max: 69, wantErr: false},
		{source: "🐞", min: 69, max: 0, wantErr: true},
		{source: "Test/n🐞", min: 0, max: 69, wantErr: false},
		{source: "Test/n🐞", min: 69, max: 0, wantErr: true},
		{source: " ", min: 0, max: 69, wantErr: false},
		{source: " ", min: 69, max: 0, wantErr: true},
		{source: "", min: 0, max: 69, wantErr: false},
		{source: "", min: 69, max: 0, wantErr: true},
	}
	for _, tt := range data {
		tt := tt
		t.Run(fmt.Sprintf("min: %d, max: %d", tt.min, tt.max), func(t *testing.T) {
			t.Parallel()
			seed := internal.SeedFrom(tt.source)
			got, err := seed.RandomInt(tt.min, tt.max)
			if err != nil && !tt.wantErr {
				t.Errorf("expected no error, got: %v", err)
				return
			}
			if !tt.wantErr && (got < tt.min || got > tt.max) {
				t.Errorf("got: %d; want: min: %d, max: %d", got, tt.min, tt.max)
			}
		})
	}
}
