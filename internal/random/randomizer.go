// Package random contains core and general purpose functionalities for the randomizer.
package random

// A Randomizer implements the logic to randomize a certain game entity.
type Randomizer interface {
	Randomize() error
}
