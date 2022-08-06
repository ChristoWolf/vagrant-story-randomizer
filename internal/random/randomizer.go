package random

// A Randomizer implements the logic to randomize a certain game entity.
type Randomizer interface {
	Randomize() error
}
