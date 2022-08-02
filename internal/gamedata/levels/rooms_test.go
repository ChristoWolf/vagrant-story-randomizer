// Package levels_test aims to cover the levels_test package.
package levels_test

import (
	"fmt"
	"testing"

	"github.com/christowolf/vagrant-story-randomizer/internal/gamedata/levels"
)

// TestZoneFileNameValid tests the ZoneFileName function
// for valid inputs.
func TestZoneFileNameValid(t *testing.T) {
	t.Parallel()
	data := []struct {
		id       int
		wantName string
	}{
		{9, "ZONE009.ZND"},
		{23, "ZONE023.ZND"},
	}
	for _, row := range data {
		row := row
		t.Run(fmt.Sprintf("%d", row.id), func(t *testing.T) {
			t.Parallel()
			gotName, err := levels.ZoneFileName(row.id)
			if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if gotName != row.wantName {
				t.Errorf("got: %s, want: %s", gotName, row.wantName)
			}
		})
	}
}

// TestZoneFileNameInvalid tests the ZoneFileName function
// for invalid inputs.
func TestZoneFileNameInvalid(t *testing.T) {
	t.Parallel()
	data := []struct {
		id int
	}{
		{1},
		{74},
		{-1},
		{0},
		{345},
		{6789},
	}
	for _, row := range data {
		row := row
		t.Run(fmt.Sprintf("%d", row.id), func(t *testing.T) {
			t.Parallel()
			gotName, err := levels.ZoneFileName(row.id)
			if err == nil {
				t.Errorf("expected an error for id: %d", row.id)
			}
			if gotName != "" {
				t.Errorf("expected empty name, got: %s", gotName)
			}
		})
	}
}

// TestMapFileName tests the MapFileName method
// for arbitrary inputs.
func TestMapFileName(t *testing.T) {
	t.Parallel()
	data := []struct {
		room     levels.Room
		wantName string
	}{
		{levels.Room{Id: 9}, "MAP009.MPD"},
		{levels.Room{Id: 283}, "MAP283.MPD"},
		{levels.Room{Id: -1}, "MAP-01.MPD"}, // Invalid, but still allowed.
	}
	for _, row := range data {
		row := row
		t.Run(fmt.Sprintf("%d", row.room.Id), func(t *testing.T) {
			t.Parallel()
			gotName, err := row.room.MapFileName()
			if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
			if gotName != row.wantName {
				t.Errorf("got: %s, want: %s", gotName, row.wantName)
			}
		})
	}
}
