package random

import (
	"errors"
	"fmt"
)

// An Id represents the ID of an in-game entity. Its value is the integer representation of its original hex value.
type Id struct {
	Value int
}

// minValue is 1, as 0 is not a valid ID.
const minValue = 1

// maxValue is 511, as the maximum ID currently supported by the randomizer in hexadecimal form is 01FF which is listed in misc items:
// https://datacrystal.romhacking.net/wiki/Vagrant_Story:misc_items_list
// Converted to an integer 01FF is 511.
const maxValue = 511

// NewId creates a new Id with the given integer as its value. If the given integer is not in the range of valid IDs then an error will be returned.
func NewId(source int) (Id, error) {
	if source < minValue || source > maxValue {
		return Id{}, fmt.Errorf("source must be between the range of %d and %d (inclusive)", minValue, maxValue)
	}
	return Id{source}, nil
}

// RandomId creates a new random Id using a certain Seed within the given Id range (inclusive).
func RandomId(seed Seed, min, max Id) (Id, error) {
	minValue := min.Value
	maxValue := max.Value
	if minValue > maxValue {
		return Id{}, errors.New("min is greater than max")
	}
	return NewId(seed.Rand().Intn(maxValue-minValue+1) + minValue)
}
