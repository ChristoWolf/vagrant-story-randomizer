package random

import (
	"hash/crc64"
	"math/rand"
	"time"
)

// A Seed represents the seed of a whole randomizer process to guarantee reproducible randomization results.
type Seed struct {
	value int64
	rand  *rand.Rand
}

// NewSeed creates a new Seed based on the current time.
func NewSeed() Seed {
	return newSeed(time.Now().UnixNano())
}

// SeedFrom creates a new Seed based on the given string's checksum.
func SeedFrom(source string) Seed {
	return newSeed(int64(crc64.Checksum([]byte(source), crc64.MakeTable(crc64.ISO))))
}

func newSeed(source int64) Seed {
	return Seed{source, rand.New(rand.NewSource(source))}
}
