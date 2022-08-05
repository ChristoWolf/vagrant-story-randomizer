package internal

import (
	"errors"
	"hash/crc64"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

type Seed struct {
	value int64
	rand  *rand.Rand
}

func NewSeed() Seed {
	return newSeed(time.Now().UnixNano())
}

func SeedFrom(source string) Seed {
	return newSeed(int64(crc64.Checksum([]byte(source), crc64.MakeTable(crc64.ISO))))
}

func newSeed(source int64) Seed {
	return Seed{source, rand.New(rand.NewSource(source))}
}

func (s Seed) RandomInt(min, max int) (int, error) {
	if min > max {
		return 0, errors.New("min is greater than max")
	}
	scaled := max - min + 1
	if scaled <= 0 {
		maxBig := big.NewInt(int64(max))
		minBig := big.NewInt(int64(min))
		var scaled big.Int
		scaled.Add(scaled.Sub(maxBig, minBig), big.NewInt(1))
		scaledInt, err := strconv.ParseInt(scaled.String(), 10, 64)
		if err != nil {
			return 0, errors.New("randomization failed for large range")
		}
		return int(scaledInt) + min, nil
	}
	return rand.Intn(max-min+1) + min, nil
}
