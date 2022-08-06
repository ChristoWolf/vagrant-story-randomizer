package random_test

import (
	"github.com/christowolf/vagrant-story-randomizer/internal/random"
	"math"
	"reflect"
	"testing"
)

func TestNewId(t *testing.T) {
	data := []struct {
		name    string
		source  int
		want    random.Id
		wantErr bool
	}{
		{name: "source 1 is valid", source: 1, want: random.Id{Value: 1}, wantErr: false},
		{name: "source 511 is valid", source: 511, want: random.Id{Value: 511}, wantErr: false},
		{name: "source 69 is valid", source: 69, want: random.Id{Value: 69}, wantErr: false},
		{name: "source 0 is not valid", source: 0, want: random.Id{}, wantErr: true},
		{name: "source 512 is not valid", source: 512, want: random.Id{}, wantErr: true},
		{name: "negative source is not valid", source: -1, want: random.Id{}, wantErr: true},
		{name: "min int source is not valid", source: math.MinInt, want: random.Id{}, wantErr: true},
		{name: "max int source is not valid", source: math.MaxInt, want: random.Id{}, wantErr: true},
	}
	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			got, err := random.NewId(tt.source)
			if err != nil && !tt.wantErr {
				t.Errorf("expected no error, got: %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestRandomId(t *testing.T) {
	data := []struct {
		name    string
		seed    random.Seed
		min     random.Id
		max     random.Id
		wantErr bool
	}{
		{name: "min is smaller than max", seed: random.NewSeed(), min: random.Id{Value: 1}, max: random.Id{Value: 69}, wantErr: false},
		{name: "min equals max", seed: random.NewSeed(), min: random.Id{Value: 69}, max: random.Id{Value: 69}, wantErr: false},
		{name: "min is greater than max", seed: random.NewSeed(), min: random.Id{Value: 69}, max: random.Id{Value: 1}, wantErr: true},
	}
	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			got, err := random.RandomId(tt.seed, tt.min, tt.max)
			if err != nil && !tt.wantErr {
				t.Errorf("expected no error, got: %v", err)
				return
			}
			if !tt.wantErr && (got.Value < tt.min.Value || got.Value > tt.max.Value) {
				t.Errorf("got: %d; want: min: %d, max: %d", got, tt.min, tt.max)
			}
		})
	}
}
